package hyperledgerfabric

import (
	"github.com/TIBCOSoftware/dovetail-go-lib/runtime/services"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type HyperledgerFabricContainerService struct {
	stub        shim.ChaincodeStubInterface
	dataService *HyperledgerFabricDataService
	evtService  *HyperledgerFabricEventService
	logService  *HyperledgerFabricLogService
}

func NewHyperledgerFabricContainerService(stub shim.ChaincodeStubInterface, loggerName string) *HyperledgerFabricContainerService {
	ctnr := &HyperledgerFabricContainerService{stub: stub}
	ctnr.dataService = NewHyperledgerFabricDataService(stub)
	ctnr.evtService = NewHyperledgerFabricEventService(stub)
	ctnr.logService = NewHyperledgerFabricLogService(stub, loggerName)
	return ctnr
}

func (ctnr *HyperledgerFabricContainerService) GetDataService() services.DataService {
	return ctnr.dataService
}
func (ctnr *HyperledgerFabricContainerService) GetEventService() services.EventService {
	return ctnr.evtService
}
func (ctnr *HyperledgerFabricContainerService) GetLogService() services.LogService {
	return ctnr.logService
}
