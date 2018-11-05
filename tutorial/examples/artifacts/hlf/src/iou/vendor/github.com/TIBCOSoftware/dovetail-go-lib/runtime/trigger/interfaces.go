package trigger

import (
	"github.com/TIBCOSoftware/dovetail-go-lib/runtime/services"
	"github.com/TIBCOSoftware/dovetail-go-lib/runtime/transaction"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

type SmartContractTrigger interface {
	trigger.Trigger
	trigger.Initializable
	Invoke(stub services.ContainerService, txn transaction.TransactionService) (status bool, data interface{}, err error)
}
