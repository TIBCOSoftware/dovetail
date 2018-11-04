package utils

import (
	"encoding/json"
	"strconv"
	"strings"
)

type outputSchema struct {
	Metadata json.RawMessage `json:"description"`
}

func ParseSchemaCollections(schema string) (map[string]map[string]interface{}, error) {
	schemametadata := make(map[string]map[string]interface{})
	allschemas := struct{ schemas [][]string }{}
	json.Unmarshal([]byte(schema), &allschemas.schemas)

	for _, value := range allschemas.schemas {
		metadata, err := ParseSchemaMetadata(value[1])
		if err != nil {
			return nil, err
		}
		schemametadata[value[0]] = metadata
	}

	return schemametadata, nil
}
func ParseSchemaMetadata(schema string) (map[string]interface{}, error) {
	//description field has metadata
	output := outputSchema{}
	json.Unmarshal([]byte(schema), &output)

	strmeta, _ := strconv.Unquote(string(output.Metadata))
	strreader := strings.NewReader(strmeta)
	decoder := json.NewDecoder(strreader)
	attrs := make(map[string]interface{})
	err := decoder.Decode(&attrs)
	return attrs, err
}

func ExtractDecoratorArgsFromSchema(decorator string, schema string) ([]interface{}, error) {
	parsedschema, err := ParseSchemaMetadata(schema)
	if err != nil {
		return nil, err
	}

	args := GetDecoratorArgs(decorator, parsedschema)
	return args, nil
}

func GetDecoratorArgs(decoratorName string, schema map[string]interface{}) []interface{} {

	metadata := schema["metadata"].(map[string]interface{})
	for _, decoratorInf := range metadata["decorators"].([]interface{}) {
		decorator := decoratorInf.(map[string]interface{})

		if decorator["name"].(string) == decoratorName {
			return decorator["args"].([]interface{})
		}
	}

	return nil
}

func GetDecoratorsByPrefix(decoratorName string, schema map[string]interface{}) []map[string]interface{} {
	decorators := make([]map[string]interface{}, 0)
	metadata := schema["metadata"].(map[string]interface{})
	for _, decoratorInf := range metadata["decorators"].([]interface{}) {
		decorator := decoratorInf.(map[string]interface{})

		if strings.HasPrefix(decorator["name"].(string), decoratorName) {
			decorators = append(decorators, decorator)
		}
	}

	return decorators
}
