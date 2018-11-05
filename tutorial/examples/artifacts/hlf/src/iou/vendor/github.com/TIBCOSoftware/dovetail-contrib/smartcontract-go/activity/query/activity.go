package query

// Imports
import (
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/utils"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Constants
const (
	ivAssetName = "assetName"
	ivData      = "queryString"
	ivParams    = "input"
	ivStub      = "containerServiceStub"
	ovOutput    = "output"
)

var logger = shim.NewLogger("FlowCC")

// describes the metadata of the activity as found in the activity.json file
type LedgerActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &LedgerActivity{metadata: metadata}
}

func (a *LedgerActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval executes the activity
func (a *LedgerActivity) Eval(context activity.Context) (done bool, err error) {
	stub, err := utils.GetContainerStub(context)
	if err != nil {
		return false, err
	}

	queryString, ok := context.GetInput(ivData).(string)
	if !ok {
		return false, fmt.Errorf("query string is not initialized")
	}

	stub.GetLogService().Debug("query string = " + queryString)

	params := context.GetInput(ivParams)
	var parameters []interface{}
	if params != nil {
		parameters, err = utils.GetInputData(params.(*data.ComplexObject), false)
		if err != nil {
			return false, err
		}
	}

	if parameters != nil && len(parameters) > 0 {
		inputParam := parameters[0].(map[string]interface{})
		for n, p := range inputParam {
			queryString = strings.Replace(queryString, "_$"+n, fmt.Sprintf("%v", p), -1)
		}
	}

	result, err := stub.GetDataService().QueryState(queryString)
	if err != nil {
		return false, err
	}

	output, err := utils.ParseRecords(result)
	if err != nil {
		return false, err
	}

	complexOutput := &data.ComplexObject{}
	complexOutput.Value = output
	context.SetOutput(ovOutput, complexOutput)

	return true, nil
}
