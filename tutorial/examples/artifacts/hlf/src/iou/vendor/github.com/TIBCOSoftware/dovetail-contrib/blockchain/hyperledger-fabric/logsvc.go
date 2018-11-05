package hyperledgerfabric

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type HyperledgerFabricLogService struct {
	stub   shim.ChaincodeStubInterface
	logger *shim.ChaincodeLogger
}

func NewHyperledgerFabricLogService(stub shim.ChaincodeStubInterface, loggerName string) *HyperledgerFabricLogService {
	return &HyperledgerFabricLogService{stub: stub, logger: shim.NewLogger(loggerName)}
}
func (logsvc *HyperledgerFabricLogService) Debug(msg string) {
	logsvc.logger.Debug(msg)
}
func (logsvc *HyperledgerFabricLogService) Info(msg string) {
	logsvc.logger.Info(msg)
}
func (logsvc *HyperledgerFabricLogService) Warning(msg string) {
	logsvc.logger.Warning(msg)
}
func (logsvc *HyperledgerFabricLogService) Error(errCode string, msg string, err error) {
	logsvc.logger.Error(errCode, msg, err)
}
