package hyperledgerfabric

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func createCompositeKey(asset map[string]interface{}, assetName string, assetKey string, stub shim.ChaincodeStubInterface) (key string, keyref string, err error) {
	key = ""
	keys := strings.Split(assetKey, ",")
	keyvalues := make([]string, 0)

	for _, k := range keys {
		k = strings.TrimSpace(k)
		id, ok := asset[k]

		if ok {
			identifierID := fmt.Sprintf("%v", id)
			keyvalues = append(keyvalues, identifierID)
		} else {
			err = fmt.Errorf("not all asset identifier(s) %s are mapped", assetKey)
			break
		}
	}

	if len(keys) == len(keyvalues) {
		key, err = stub.CreateCompositeKey(assetName, keyvalues)
		if err != nil {
			err = fmt.Errorf("error creating composite key for %s with attribute %s, err: %v", assetName, assetKey, err)
		} else {
			keyref, err = stub.CreateCompositeKey(getSecondaryCompositeNS(assetName), keyvalues)
		}
	}

	return key, keyref, err
}

func createPartialCompositeKey(asset map[string]interface{}, assetKey string, stub shim.ChaincodeStubInterface) (keyvalues []string, err error) {
	keys := strings.Split(assetKey, ",")
	keyvalues = make([]string, 0)

	for idx, k := range keys {
		k = strings.TrimSpace(k)
		id, ok := asset[k]

		if ok {
			identifierID := fmt.Sprintf("%v", id)
			if len(keyvalues) < idx {
				err = fmt.Errorf("error to create partial composite lookup key, composite key is %s, key %s is not mapped ", assetKey, keys[idx-1])
				break
			}
			keyvalues = append(keyvalues, identifierID)
		}
	}

	return
}

func getSecondaryCompositeNS(assetName string) string {
	return assetName + "_refs"
}

func getRecord(assetName string, identifier string, input map[string]interface{}, stub shim.ChaincodeStubInterface) ([]byte, error) {
	key, _, err := createCompositeKey(input, assetName, identifier, stub)
	if err != nil {
		return nil, err
	}

	rawvalue, err := stub.GetState(key)
	if err != nil {
		return nil, fmt.Errorf("error GetState with key %s from ledger", key)
	}

	if rawvalue != nil && len(rawvalue) > 0 {
		return rawvalue, nil
	} else {
		return nil, nil
	}
}

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
		compkey, _, err := createCompositeKey(record, getSecondaryCompositeNS(assetName), key, stub)
		if err != nil {
			return nil, err
		}

		compkeys = append(compkeys, compkey)
	}

	return compkeys, nil
}
func parseRecord(record []byte) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := json.Unmarshal(record, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func deleteRecord(assetName string, identifier string, input map[string]interface{}, compositeKeys interface{}, stub shim.ChaincodeStubInterface) ([]byte, error) {
	key, keyref, err := createCompositeKey(input, assetName, identifier, stub)
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
}

func putRecord(assetName string, identifier string, input map[string]interface{}, compositeKeys interface{}, stub shim.ChaincodeStubInterface) error {
	bv, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("error martial value to json for %s with attribute %s, err: %v", assetName, identifier, err)
	}

	key, keyref, err := createCompositeKey(input, assetName, identifier, stub)
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
}

func lookupRecords(assetName string, identifier string, lookupKey interface{}, input map[string]interface{}, stub shim.ChaincodeStubInterface) ([][]byte, error) {
	result := make([][]byte, 0)
	keyvalues, err := createPartialCompositeKey(input, lookupKey.(string), stub)
	if err != nil {
		return nil, err
	}

	var sameAsIdentifier = true
	namespace := assetName
	if identifier != lookupKey.(string) {
		sameAsIdentifier = false
		namespace = getSecondaryCompositeNS(assetName)
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
}

func (svc *HyperledgerFabricDataService) LookupState(assetName string, assetKey string, lookupKey string, lkupKeyValue map[string]interface{}) ([][]byte, error) {
	result := make([][]byte, 0)
	keyvalues, err := createPartialCompositeKey(lkupKeyValue, lookupKey, svc.stub)
	if err != nil {
		return nil, err
	}

	var sameAsIdentifier = true
	namespace := assetName
	if assetKey != lookupKey {
		sameAsIdentifier = false
		namespace = getSecondaryCompositeNS(assetName)
	}

	iterator, err := svc.stub.GetStateByPartialCompositeKey(namespace, keyvalues)
	if err != nil {
		return nil, fmt.Errorf("error GetStateByPartialCompositeKey with key %s from ledger", lookupKey)
	}
	defer iterator.Close()

	for iterator.HasNext() {
		kv, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("error iterating results from GetStateByPartialCompositeKey with key %s from ledger", lookupKey)
		}

		if sameAsIdentifier {
			result = append(result, kv.GetValue())
		} else {
			primary := string(kv.GetValue())
			v, err := svc.stub.GetState(primary)
			if err != nil {
				return nil, err
			}

			result = append(result, v)
		}
	}

	return result, nil

}
