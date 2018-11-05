package transaction

import (
	"context"
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/dovetail-go-lib/runtime/services"

	"github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/activity/txnreply"
	"github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/utils"
	rttxn "github.com/TIBCOSoftware/dovetail-go-lib/runtime/transaction"
	fldata "github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Trigger must define a struct
type TXNTrigger struct {
	// Trigger Metadata
	metadata *trigger.Metadata

	// Trigger configuration
	config *trigger.Config

	handlers map[string]*TXNHandler
}

type TXNHandler struct {
	metadtaSchema string
	outputs       []rttxn.TxnInputAttribute
	handler       *trigger.Handler
	acl           rttxn.TxnACL
}

// Trigger must define a factory
type TXNTriggerFactory struct {
	// Trigger Metadata
	metadata *trigger.Metadata
}

// NewFactory create a new Trigger factory
// Trigger must define this function
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &TXNTriggerFactory{metadata: md}
}

// Creates a new trigger instance for a given id
// Trigger must define this method
func (t *TXNTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	txntrigger := &TXNTrigger{metadata: t.metadata, config: config, handlers: make(map[string]*TXNHandler)}
	//setting schemas has all the generated json schema from connector
	allschemas, _ := utils.ParseSchemaCollections(config.GetSetting("schemas"))

	for _, handler := range config.Handlers {
		txn := handler.GetSetting("transaction")
		txnHandler := parseHandlerMetadata(handler, allschemas)
		txntrigger.handlers[txn] = txnHandler

		//also store handler without namespace
		tokens := strings.Split(txn, ".")
		txntrigger.handlers[tokens[len(tokens)-1]] = txnHandler
	}

	return txntrigger
}

// Returns trigger metadata
// Trigger must define this method
func (t *TXNTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

func (t *TXNTrigger) Initialize(ctx trigger.InitContext) error {
	if len(ctx.GetHandlers()) == 0 {
		return fmt.Errorf("no Handlers found for trigger '%s'", t.config.Id)
	}

	for _, handler := range ctx.GetHandlers() {
		txn, _ := handler.GetSetting("transaction")
		t.handlers[txn.(string)].handler = handler
	}

	return nil
}

// Start trigger. Start will be called once engine is started successfully.
func (t *TXNTrigger) Start() error {
	return nil
}

// Stop trigger. Stop will be called in case engine is gracefully stopped.
func (t *TXNTrigger) Stop() error {
	return nil
}

//Trigger must implement this function from dovetail-go-lib/runtime/trigger SmartContractTrigger interface
func (t *TXNTrigger) Invoke(stub services.ContainerService, txn rttxn.TransactionService) (status bool, data interface{}, err error) {
	logger.Debug("Enter trigger Invoke...")
	defer logger.Debug("Exit trigger invoke")
	handlerinfo := t.handlers[txn.GetTransactionName()]
	if handlerinfo == nil {
		return false, nil, fmt.Errorf("handler for transaction %s is not found", txn.GetTransactionName())
	}

	txnInput, err := txn.ResolveTransactionInput(handlerinfo.outputs)
	if err != nil {
		return false, nil, err
	}

	if txn.TransactionSecuritySupported() {
		authorized, err := aclCheck(txn, txnInput, handlerinfo.acl)
		if err != nil {
			return false, nil, err
		}

		if !authorized {
			return false, nil, fmt.Errorf("Transaction creator is not authorized to initiate the transaction %s", txn)
		}
	}

	triggerdata := make(map[string]interface{})
	triggerdata["transactionInput"] = &fldata.ComplexObject{Metadata: handlerinfo.metadtaSchema, Value: txnInput}
	triggerdata["containerServiceStub"] = stub

	results, err := handlerinfo.handler.Handle(context.Background(), triggerdata)
	if err != nil {
		return false, nil, err
	}

	//reply data
	reply := results["data"].Value().(txnreply.Reply)
	switch reply.Status {
	case txnreply.SUCCESS:
		return true, nil, nil
	case txnreply.SUCCESS_WITH_DATA:
		return true, reply.Payload, nil
	default:
		return false, nil, fmt.Errorf(reply.Message)
	}
}

func aclCheck(txn rttxn.TransactionService, txInput map[string]interface{}, acl rttxn.TxnACL) (bool, error) {
	creator, err := txn.GetTransactionInitiator()
	logger.Debugf("Caller identity=%s\n", creator)
	if err != nil {
		return false, err
	}

	if len(acl.AuthorizedParty) > 0 {
		for _, participant := range acl.AuthorizedParty {
			if strings.TrimSpace(participant) == "*" {
				return true, nil
			}

			//$tx.path.to.party
			value, err := utils.FindValueInMap(txInput, participant)
			if err != nil {
				return false, fmt.Errorf("invalid ACL format for InitiatedBy, expected $tx.path.to.party")
			}

			id := value.(string)

			if id == creator {
				return true, nil
			} else {
				return false, nil
			}
		}
		return false, nil
	} else {
		return isAuthorized(txn, acl)
	}
}

func isAuthorized(txn rttxn.TransactionService, acl rttxn.TxnACL) (bool, error) {
	for attr, condition := range acl.Conditions {
		value, found, err := txn.GetInitiatorCertAttribute(attr)
		if err != nil {
			return false, err
		}

		if !found || value != condition {
			return false, nil
		}
	}

	return true, nil
}
