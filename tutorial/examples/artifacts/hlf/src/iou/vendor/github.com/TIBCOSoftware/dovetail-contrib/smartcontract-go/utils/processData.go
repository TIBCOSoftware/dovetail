package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	dtsvc "github.com/TIBCOSoftware/dovetail-go-lib/runtime/services"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

func FindValueInMap(input map[string]interface{}, attrpath string) (interface{}, error) {
	paths := strings.Split(attrpath, ".")
	if len(paths) < 2 {
		return nil, fmt.Errorf("invalid format, expected $tx.path.to.attribute")
	}

	value := input[paths[1]]
	if len(paths) >= 2 {
		for idx := 2; idx < len(paths); idx++ {
			valuem := value.(map[string]interface{})
			value = valuem[paths[idx]]
		}
	}
	return value, nil
}

func GetInputData(assetValue *data.ComplexObject, isArray bool) ([]interface{}, error) {
	var rawInput interface{}
	assets := make([]interface{}, 0)

	if assetValue == nil {
		return nil, nil
	}
	//fmt.Printf("input=%v\n, type=%s\n", assetValue.Value, reflect.TypeOf(assetValue.Value))
	switch t := assetValue.Value.(type) {
	case string:
		if t != "" && t != "{}" {
			err := json.Unmarshal([]byte(t), &rawInput)
			if err != nil {
				return nil, err
			}

			if isArray {
				assets = rawInput.([]interface{})
			} else {
				assets = append(assets, rawInput.(interface{}))
			}
		}
		break
	case []string:
		for _, v := range t {
			err := json.Unmarshal([]byte(v), &rawInput)
			if err != nil {
				return nil, err
			}

			assets = append(assets, rawInput.(interface{}))
		}
	case map[string]interface{}:
		assets = append(assets, t)
		break
	case []map[string]interface{}:
		for _, v := range t {
			assets = append(assets, v)
		}
		break
	case []interface{}:
		for _, v := range t {
			iv, err := convertToMap(v)
			if err != nil {
				return nil, err
			}
			assets = append(assets, iv)
		}
		break
	case interface{}:
		iv, err := convertToMap(t)
		if err != nil {
			return nil, err
		}
		assets = append(assets, iv)
		break
	}
	return assets, nil
}

func convertToMap(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case map[string]interface{}:
		return v, nil
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		var m interface{}
		err = json.Unmarshal(b, &m)
		if err != nil {
			return nil, err
		}
		return m, nil
	}
}

func ParseRecord(record []byte) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := json.Unmarshal(record, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ParseRecords(records [][]byte) ([]interface{}, error) {
	output := make([]interface{}, 0)
	for _, v := range records {
		m, err := ParseRecord(v)
		if err != nil {
			return output, err
		}
		output = append(output, m)
	}

	return output, nil
}

func GetContainerStub(context activity.Context) (dtsvc.ContainerService, error) {
	stubobj := context.GetInput("containerServiceStub")
	if stubobj == nil {
		return nil, fmt.Errorf("containerServiceStub is not initialized")
	}

	stub, ok := interface{}(stubobj).(dtsvc.ContainerService)
	if !ok {
		return nil, fmt.Errorf("containerServiceStub is not instance of ContainerService")
	}

	return stub, nil
}

func IsPrimitive(datatype string) bool {
	result := false
	dt := strings.ToLower(datatype)
	switch dt {
	case "string", "integer", "long", "boolean", "double", "datetime":
		result = true
		break
	}
	return result
}
