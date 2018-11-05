package hyperledgerfabric

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	//"github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/utils"
	"github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/utils"
	txn "github.com/TIBCOSoftware/dovetail-go-lib/runtime/transaction"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type HyperledgerFabricTransactionService struct {
	stub            shim.ChaincodeStubInterface
	args            []string
	name            string
	securityEnabled bool
}

func NewHyperledgerFabricTransactionService(stub shim.ChaincodeStubInterface, args []string, name string, enableSecurity bool) *HyperledgerFabricTransactionService {
	return &HyperledgerFabricTransactionService{stub: stub, args: args, name: name, securityEnabled: enableSecurity}
}

func (svc *HyperledgerFabricTransactionService) GetTransactionName() string {
	return svc.name
}

func (svc *HyperledgerFabricTransactionService) TransactionSecuritySupported() bool {
	return svc.securityEnabled
}

func (svc *HyperledgerFabricTransactionService) ResolveTransactionInput(txnInputsMetadata []txn.TxnInputAttribute) (map[string]interface{}, error) {
	txnInput := make(map[string]interface{})
	//add transactionId and timestamp,should be last 2 argments
	svc.args = append(svc.args, svc.stub.GetTxID())
	txntime, err := svc.stub.GetTxTimestamp()
	if err != nil {
		return txnInput, fmt.Errorf("Error retrieving transaction timestamp: %v", err)
	}
	txnstr := time.Unix(txntime.Seconds, int64(txntime.Nanos)).UTC().Format("2006-01-02T15:04:05.00000-0700")
	svc.args = append(svc.args, txnstr)

	if len(txnInputsMetadata) != len(svc.args) {
		return txnInput, fmt.Errorf("Expected %v arguments: %v, only found %v from input: %v", len(txnInputsMetadata), txnInputsMetadata, len(svc.args), svc.args)
	}

	for idx, arg := range svc.args {
		attr := txnInputsMetadata[idx]
		if attr.IsAssetRef {
			//resolve asset reference
			var assetkeyvalues interface{}
			err := json.Unmarshal([]byte(arg), &assetkeyvalues)
			if err != nil {
				return txnInput, err
			}
			if attr.IsArray {
				input := make([]map[string]interface{}, 0)
				keys := assetkeyvalues.([]interface{})
				for _, k := range keys {
					rc, err := resolveRecord(attr.Name, attr.AssetName, attr.Identifiers, k.(map[string]interface{}), svc.stub)
					if err != nil {
						return txnInput, err
					}
					input = append(input, rc)
				}
				txnInput[attr.Name] = input
			} else {
				rc, err := resolveRecord(attr.Name, attr.AssetName, attr.Identifiers, assetkeyvalues.(map[string]interface{}), svc.stub)
				if err != nil {
					return txnInput, err
				}
				txnInput[attr.Name] = rc
			}
		} else {
			if attr.IsArray || (!utils.IsPrimitive(attr.DataType) && !attr.IsParticipant) {
				var input interface{}
				err := json.Unmarshal([]byte(arg), &input)
				if err != nil {
					return txnInput, err
				}
				txnInput[attr.Name] = input
			} else {
				input, err := parsePrimitiveType(attr, arg)
				if err != nil {
					return txnInput, err
				}
				txnInput[attr.Name] = input
			}
		}
	}

	return txnInput, nil
}

func (svc *HyperledgerFabricTransactionService) GetTransactionInitiator() (string, error) {

	return cid.GetMSPID(svc.stub)
}

func (svc *HyperledgerFabricTransactionService) GetInitiatorCertAttribute(attr string) (value string, found bool, err error) {
	return cid.GetAttributeValue(svc.stub, attr)
}

func resolveRecord(attrName, assetName, identifiers string, values map[string]interface{}, stub shim.ChaincodeStubInterface) (map[string]interface{}, error) {
	logger.Debugf("trigger::resolveRecord - attrName=%s, assetName=%s, identifiers=%s, values=%v", attrName, assetName, identifiers, values)
	rc, err := getRecord(assetName, identifiers, values, stub)
	if err != nil {
		return nil, err
	}

	if rc == nil {
		return nil, fmt.Errorf("Asset reference %s cann't be resolved", attrName)
	}

	input := make(map[string]interface{})
	err = json.Unmarshal(rc, &input)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func parsePrimitiveType(attr txn.TxnInputAttribute, arg string) (interface{}, error) {
	switch attr.DataType {
	case "Integer", "Long":
		iv, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		return iv, nil
	case "Boolean":
		bv, err := strconv.ParseBool(arg)
		if err != nil {
			return nil, err
		}
		return bv, nil
	default:
		return arg, nil
	}

}
