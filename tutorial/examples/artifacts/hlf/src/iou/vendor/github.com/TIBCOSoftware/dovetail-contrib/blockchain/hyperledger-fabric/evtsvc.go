package hyperledgerfabric

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type HyperledgerFabricEventService struct {
	stub shim.ChaincodeStubInterface
}

func NewHyperledgerFabricEventService(stub shim.ChaincodeStubInterface) *HyperledgerFabricEventService {
	return &HyperledgerFabricEventService{stub: stub}
}

func (evt *HyperledgerFabricEventService) Publish(evtName, metadata string, evtPayload []byte) error {
	if metadata != "" {
		evtName = evtName + "#" + metadata
	}
	evt.stub.SetEvent(evtName, evtPayload)
	return nil
}
