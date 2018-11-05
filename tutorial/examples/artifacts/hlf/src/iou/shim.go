package main

import (
	"fmt"
	hlf "github.com/TIBCOSoftware/dovetail-contrib/blockchain/hyperledger-fabric"
	dttrigger "github.com/TIBCOSoftware/dovetail-go-lib/runtime/trigger"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type IOUChainCode struct {
	ContractName string
	TxnTrigger   dttrigger.SmartContractTrigger
}

func (f *IOUChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (f *IOUChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	hlfctnr := hlf.NewHyperledgerFabricContainerService(stub, "IOU")
	logger := hlfctnr.GetLogService()
	fn, args := stub.GetFunctionAndParameters()

	logger.Info(fmt.Sprintf("fn=%s, args=%s\n", fn, args[0]))

	txn := hlf.NewHyperledgerFabricTransactionService(stub, args, fn, enableSecurity)
	status, data, err := f.TxnTrigger.Invoke(hlfctnr, txn)
	logger.Info(fmt.Sprintf("status=%v, data=%v, err=%v\n", status, data, err))
	if err == nil {
		if data == nil {
			return shim.Success(nil)
		} else {
			return shim.Success([]byte(data.(string)))
		}
	} else {
		return shim.Error(fmt.Sprintf("failed to execute transaction: %s, error:%v\n", fn, err))
	}
}

var flowcc = IOUChainCode{ContractName: "IOU"}
var logger = shim.NewLogger("IOU_InitEngine")

func main() {
	if err := shim.Start(&flowcc); err != nil {
		logger.Error(fmt.Sprintf("Error starting Chain code %s: %s", flowcc.ContractName, err))
	}
}