package logger

// Imports
import (
	"github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/utils"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// Constants
const (
	ivLogLevel = "logLevel"
	ivMessage  = "message"
	ivErrcode  = "errorCode"
	ivStub     = "containerServiceStub"
	ivContract = "FlowCC"
)

// describes the metadata of the activity as found in the activity.json file
type LoggerActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &LoggerActivity{metadata: metadata}
}

func (a *LoggerActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *LoggerActivity) Eval(context activity.Context) (done bool, err error) {
	stub, err := utils.GetContainerStub(context)
	if err != nil {
		return false, err
	}

	appLogLevel, _ := context.GetInput(ivLogLevel).(string)
	message, _ := context.GetInput(ivMessage).(string)
	errcode, _ := context.GetInput(ivErrcode).(string)

	activityLog := stub.GetLogService()

	switch appLogLevel {
	case "INFO":
		activityLog.Info(message)
		break
	case "DEBUG":
		activityLog.Debug(message)
		break
	case "WARNING":
		activityLog.Warning(message)
		break
	case "ERROR":
		activityLog.Error(errcode, message, nil)
		break
	}

	return true, nil
}
