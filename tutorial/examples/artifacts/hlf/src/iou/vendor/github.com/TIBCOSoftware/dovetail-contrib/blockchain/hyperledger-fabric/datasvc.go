package hyperledgerfabric

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type HyperledgerFabricDataService struct {
	stub shim.ChaincodeStubInterface
}

func NewHyperledgerFabricDataService(stub shim.ChaincodeStubInterface) *HyperledgerFabricDataService {
	return &HyperledgerFabricDataService{stub: stub}
}

func (svc *HyperledgerFabricDataService) PutState(assetName string, assetKey string, assetValue map[string]interface{}, secondaryCompKeys interface{}) error {
	bv, err := json.Marshal(assetValue)
	if err != nil {
		return fmt.Errorf("error martial value to json for %s with attribute %s, err: %v", assetName, assetKey, err)
	}

	key, keyref, err := createCompositeKey(assetValue, assetName, assetKey, svc.stub)
	if err != nil {
		return err
	}

	if secondaryCompKeys != nil && secondaryCompKeys.(string) != "" {
		err = storeCompositeKeyRefs(assetName, key, keyref, secondaryCompKeys.(string), assetValue, svc.stub)

		if err != nil {
			return err
		}
	}
	err = svc.stub.PutState(key, bv)
	if err != nil {
		return err
	}

	return nil
}

func (svc *HyperledgerFabricDataService) GetState(assetName string, assetKey string, keyValue map[string]interface{}) ([]byte, error) {
	return getRecord(assetName, assetKey, keyValue, svc.stub)
}

func (svc *HyperledgerFabricDataService) DeleteState(assetName string, assetKey string, keyValue map[string]interface{}, secondaryCompKeys interface{}) ([]byte, error) {
	key, keyref, err := createCompositeKey(keyValue, assetName, assetKey, svc.stub)
	if err != nil {
		return nil, err
	}
	rawvalue, err := svc.stub.GetState(key)

	if err != nil {
		return nil, err
	}

	if rawvalue != nil && len(rawvalue) > 0 {
		err = svc.stub.DelState(key)
		if err != nil {
			return nil, err
		}

		if secondaryCompKeys != nil && secondaryCompKeys.(string) != "" {
			refs, err := svc.stub.GetState(keyref)
			if err != nil {
				return nil, err
			}
			err = delCompositeKeyRefs(keyref, refs, svc.stub)
			if err != nil {
				return nil, err
			}
		}
		return rawvalue, nil
	}
	return nil, nil
}

func (svc *HyperledgerFabricDataService) GetHistory(assetName string, assetKey string, keyValue map[string]interface{}) ([][]byte, error) {
	output := make([][]byte, 0)
	key, _, err := createCompositeKey(keyValue, assetName, assetKey, svc.stub)
	if err != nil {
		return output, err
	}

	iterator, err := svc.stub.GetHistoryForKey(key)
	if err != nil {
		return output, err
	}
	defer iterator.Close()

	for iterator.HasNext() {
		kv, err := iterator.Next()
		if err != nil {
			return output, err
		}

		output = append(output, kv.GetValue())
	}

	return output, nil
}
func (svc *HyperledgerFabricDataService) QueryState(queryString string) ([][]byte, error) {
	output := make([][]byte, 0)
	iterator, err := svc.stub.GetQueryResult(queryString)
	if err != nil {
		return output, err
	}
	defer iterator.Close()

	for iterator.HasNext() {
		kv, err := iterator.Next()
		if err != nil {
			return output, err
		}

		output = append(output, kv.GetValue())
	}
	return output, nil
}
