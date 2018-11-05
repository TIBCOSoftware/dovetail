package ledger

// Imports
import (
	"encoding/json"
	"fmt"

	"github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/utils"
	"github.com/TIBCOSoftware/flogo-lib/core/data"

	dtsvc "github.com/TIBCOSoftware/dovetail-go-lib/runtime/services"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// Constants
const (
	ivAssetName      = "assetName"
	ivData           = "input"
	ivStub           = "containerServiceStub"
	ivOperation      = "operation"
	ivAssetKey       = "identifier"
	ivAssetLookupKey = "compositeKey"
	ivCompositeKeys  = "compositeKeys"
	ivIsArray        = "isArray"
	ovOutput         = "output"
)

// describes the metadata of the activity as found in the activity.json file
type LedgerActivity struct {
	metadata *activity.Metadata
}

// NewActivity will instantiate a new LedgerActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &LedgerActivity{metadata: metadata}
}

// Metadata will return the metadata of the LedgerActivity
func (a *LedgerActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval executes the activity
func (a *LedgerActivity) Eval(context activity.Context) (done bool, err error) {
	stub, err := utils.GetContainerStub(context)
	if err != nil {
		return false, err
	}

	logger := stub.GetLogService()

	logger.Debug("Enter ledger activity...")
	isArray, ok := context.GetInput(ivIsArray).(bool)
	if !ok {
		return false, fmt.Errorf("asset name is not initialized")
	}

	assetName, ok := context.GetInput(ivAssetName).(string)
	if !ok {
		return false, fmt.Errorf("asset name is not initialized")
	}

	assetKey, ok := context.GetInput(ivAssetKey).(string)
	if !ok {
		return false, fmt.Errorf("asset key is not initialized")
	}
	operation, ok := context.GetInput(ivOperation).(string)
	if !ok {
		return false, fmt.Errorf("operation is not initialized")
	}

	inputValue, err := data.CoerceToComplexObject(context.GetInput(ivData))
	if err != nil {
		return false, fmt.Errorf("asset value is not initialized")
	}

	if inputValue == nil {
		return false, fmt.Errorf("asset value is not initialized")
	}
	inputs, err := utils.GetInputData(inputValue, isArray)
	if err != nil {
		return false, err
	}

	compositeKeys := context.GetInput(ivCompositeKeys)
	lookupKey := context.GetInput(ivAssetLookupKey)

	result := make([][]byte, 0)
	output := make([]interface{}, 0)
	complexOutput := &data.ComplexObject{}
	for _, av := range inputs {
		asset, ok := av.(map[string]interface{})
		if !ok {
			return false, fmt.Errorf("can not cast unmartialed asset value to instance of map[string]interface{}")
		}
		fmt.Printf("ledger input %v\n", asset)
		if operation == "DELETE" {
			record, err := deleteRecord(assetName, assetKey, asset, compositeKeys, stub)
			if err != nil {
				return false, err
			}

			if record == nil {
				output = append(output, asset)
			} else {
				result = append(result, record)
			}
		} else if operation == "PUT" {
			err = putRecord(assetName, assetKey, asset, compositeKeys, stub)
			if err != nil {
				return false, err
			}
			output = append(output, asset)
		} else if operation == "GET" {
			rawvalue, err := getRecord(assetName, assetKey, asset, stub)
			if err != nil {
				return false, err
			}
			if rawvalue != nil {
				result = append(result, rawvalue)
			}
		} else {
			//LOOKUP
			records, err := lookupRecords(assetName, assetKey, lookupKey.(string), asset, stub)
			if err != nil {
				return false, err
			}

			if records != nil && len(records) > 0 {
				result = append(result, records...)
			}
		}
	}

	for _, v := range result {
		m, err := utils.ParseRecord(v)
		if err != nil {
			return false, err
		}
		output = append(output, m)
	}

	if isArray || operation == "LOOKUP" {
		complexOutput.Value = output
	} else {
		if len(output) > 0 {
			complexOutput.Value = output[0]
		}
	}

	context.SetOutput(ovOutput, complexOutput)
	logger.Debug("Exit ledger activity")
	return true, nil
}

/*
func storeCompositeKeyRefs(assetName string, primaryKey string, primaryKeyRef string, compositeKeys string, record map[string]interface{}, stub shim.ChaincodeStubInterface) error {
	compkeys, err := createSecondaryCompositeKeys(assetName, compositeKeys, record, stub)
	if err != nil {
		return err
	}

	refs, err := stub.GetState(primaryKeyRef)
	if err != nil {
		return err
	}

	encodedcomkeys := &bytes.Buffer{}
	gob.NewEncoder(encodedcomkeys).Encode(compkeys)
	if refs != nil {
		if bytes.Compare(refs, encodedcomkeys.Bytes()) != 0 {

			err = delCompositeKeyRefs(primaryKeyRef, refs, stub)
			if err != nil {
				return err
			}

			return writeCompositeKeysToDB(stub, compkeys, encodedcomkeys.Bytes(), primaryKey, primaryKeyRef)
		}
	} else {
		return writeCompositeKeysToDB(stub, compkeys, encodedcomkeys.Bytes(), primaryKey, primaryKeyRef)
	}

	return nil
}

func writeCompositeKeysToDB(stub shim.ChaincodeStubInterface, keys []string, encodedKeys []byte, primaryKey string, primaryKeyRef string) error {
	for _, k := range keys {
		err := stub.PutState(k, []byte(primaryKey))
		if err != nil {
			return err
		}
	}

	err := stub.PutState(primaryKeyRef, encodedKeys)
	return err
}

func delCompositeKeyRefs(primaryKeyRef string, refs []byte, stub shim.ChaincodeStubInterface) error {
	comkeys := []string{}
	gob.NewDecoder(bytes.NewReader(refs)).Decode(&comkeys)
	for _, k := range comkeys {
		err := stub.DelState(k)
		if err != nil {
			return err
		}
	}
	return stub.DelState(primaryKeyRef)
}

func createSecondaryCompositeKeys(assetName string, keys string, record map[string]interface{}, stub shim.ChaincodeStubInterface) ([]string, error) {

	keyarray := strings.Split(keys, "|")
	compkeys := make([]string, 0)
	for _, key := range keyarray {
		compkey, _, err := hlfutils.CreateCompositeKey(record, hlfutils.GetSecondaryCompositeNS(assetName), key, stub)
		if err != nil {
			return nil, err
		}

		compkeys = append(compkeys, compkey)
	}

	return compkeys, nil
}
*/
func parseRecord(record []byte) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := json.Unmarshal(record, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func deleteRecord(assetName string, identifier string, input map[string]interface{}, compositeKeys interface{}, stub dtsvc.ContainerService) ([]byte, error) {
	datasvc := stub.GetDataService()
	return datasvc.DeleteState(assetName, identifier, input, compositeKeys)
	/*
		key, keyref, err := hlfutils.CreateCompositeKey(input, assetName, identifier, stub)
		if err != nil {
			return nil, err
		}
		rawvalue, err := stub.GetState(key)

		if err != nil {
			return nil, err
		}

		if rawvalue != nil && len(rawvalue) > 0 {
			err = stub.DelState(key)
			if err != nil {
				return nil, err
			}

			if compositeKeys != nil && compositeKeys.(string) != "" {
				refs, err := stub.GetState(keyref)
				if err != nil {
					return nil, err
				}
				err = delCompositeKeyRefs(keyref, refs, stub)
				if err != nil {
					return nil, err
				}
			}
			return rawvalue, nil
		}
		return nil, nil
	*/
}

func putRecord(assetName string, identifier string, input map[string]interface{}, compositeKeys interface{}, stub dtsvc.ContainerService) error {

	datasvc := stub.GetDataService()
	return datasvc.PutState(assetName, identifier, input, compositeKeys)
	/*
		bv, err := json.Marshal(input)
		if err != nil {
			return fmt.Errorf("error martial value to json for %s with attribute %s, err: %v", assetName, identifier, err)
		}

		key, keyref, err := hlfutils.CreateCompositeKey(input, assetName, identifier, stub)
		if err != nil {
			return err
		}

		if compositeKeys != nil && compositeKeys.(string) != "" {
			err = storeCompositeKeyRefs(assetName, key, keyref, compositeKeys.(string), input, stub)

			if err != nil {
				return err
			}
		}
		err = stub.PutState(key, bv)
		if err != nil {
			return err
		}

		return nil
	*/
}

func lookupRecords(assetName string, identifier string, lookupKey string, input map[string]interface{}, stub dtsvc.ContainerService) ([][]byte, error) {
	datasvc := stub.GetDataService()
	return datasvc.LookupState(assetName, identifier, lookupKey, input)
	/*result := make([][]byte, 0)
	keyvalues, err := hlfutils.CreatePartialCompositeKey(input, lookupKey.(string), stub)
	if err != nil {
		return nil, err
	}

	var sameAsIdentifier = true
	namespace := assetName
	if identifier != lookupKey.(string) {
		sameAsIdentifier = false
		namespace = hlfutils.GetSecondaryCompositeNS(assetName)
	}

	iterator, err := stub.GetStateByPartialCompositeKey(namespace, keyvalues)
	if err != nil {
		return nil, fmt.Errorf("error GetStateByPartialCompositeKey with key %s from ledger", lookupKey.(string))
	}
	defer iterator.Close()

	for iterator.HasNext() {
		kv, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("error iterating results from GetStateByPartialCompositeKey with key %s from ledger", lookupKey.(string))
		}

		if sameAsIdentifier {
			result = append(result, kv.GetValue())
		} else {
			primary := string(kv.GetValue())
			v, err := stub.GetState(primary)
			if err != nil {
				return nil, err
			}

			result = append(result, v)
		}
	}

	return result, nil
	*/
}

func getRecord(assetName string, identifier string, input map[string]interface{}, stub dtsvc.ContainerService) ([]byte, error) {
	datasvc := stub.GetDataService()
	return datasvc.GetState(assetName, identifier, input)
}
