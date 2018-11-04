package txnreply

// Imports
import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// Constants
const (
	ivStatus   = "status"
	ivMessage  = "message"
	ivData     = "input"
	ivUserData = "userInput"
	ivDataType = "dataType"
)

const (
	SUCCESS           = "Success"
	SUCCESS_WITH_DATA = "Success with Data"
	ERROR             = "Error with Message"
)

type Reply struct {
	Status  string
	Message string
	Payload string
}

// describes the metadata of the activity as found in the activity.json file
type TxnResponseActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &TxnResponseActivity{metadata: metadata}
}

func (a *TxnResponseActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *TxnResponseActivity) Eval(context activity.Context) (done bool, err error) {
	replyHandler := context.FlowDetails().ReplyHandler()

	//todo support replying with error
	if replyHandler == nil {
		return false, fmt.Errorf("Reply handler is not available in the flow")
	}

	code := 0
	reply := Reply{}
	err = nil
	status := context.GetInput(ivStatus).(string)
	reply.Status = strings.TrimSpace(status)
	dataType := context.GetInput(ivDataType).(string)
	if strings.Compare(reply.Status, SUCCESS_WITH_DATA) == 0 {
		rawinput := context.GetInput(ivData)
		if dataType == "User Defined..." {
			rawinput = context.GetInput(ivUserData)
		}
		resp, _ := data.CoerceToComplexObject(rawinput)
		respstr, err := json.Marshal(resp.Value)
		if err != nil {
			return false, err
		}
		reply.Payload = string(respstr)
	} else if strings.Compare(reply.Status, ERROR) == 0 {
		code = 1
		err = fmt.Errorf(context.GetInput(ivMessage).(string))
	}
	replyHandler.Reply(code, reply, err)
	logger.Debug("Exit txreply activity")
	return true, nil
}
