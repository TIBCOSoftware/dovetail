package transaction

import (
	"strings"

	"github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/utils"
	rttxn "github.com/TIBCOSoftware/dovetail-go-lib/runtime/transaction"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

func parseHandlerMetadata(handler *trigger.HandlerConfig, allschemas map[string]map[string]interface{}) *TXNHandler {

	handlerinfo := TXNHandler{}

	jsonschema := handler.Outputs["transactionInput"].(map[string]interface{})["metadata"]
	handlerinfo.metadtaSchema = jsonschema.(string)

	schema, _ := utils.ParseSchemaMetadata(jsonschema.(string))

	outputattrs := make([]rttxn.TxnInputAttribute, 0)
	for _, attrInf := range schema["attributes"].([]interface{}) {
		attr := attrInf.(map[string]interface{})
		dttype := attr["type"].(string)
		isAssetRef := false
		isArray := false
		assetName := ""
		identifiers := ""
		isParticipant := false
		if !utils.IsPrimitive(dttype) {
			schema := allschemas[dttype]
			meta := schema["metadata"].(map[string]interface{})
			if isRef, ok := attr["isRef"].(bool); ok && isRef {
				//check if it is asset reference
				if meta["type"].(string) == "Asset" {
					isAssetRef = true
					assetName = dttype
					identifiers = meta["identifiedBy"].(string)
				}
			}

			if meta["type"].(string) == "Participant" {
				isParticipant = true
			}
		}

		if ar, ok := attr["isArray"]; ok {
			isArray = ar.(bool)
		}
		outputattrs = append(outputattrs, rttxn.TxnInputAttribute{Name: attr["name"].(string), DataType: dttype, IsAssetRef: isAssetRef, AssetName: assetName, Identifiers: identifiers, IsArray: isArray, IsParticipant: isParticipant})
	}

	//access control
	args := utils.GetDecoratorArgs("InitiatedBy", schema)
	acl := rttxn.TxnACL{}

	if args == nil || len(args) == 0 {
		acl.AuthorizedParty = make([]string, 1)
		acl.AuthorizedParty[0] = "*"
	} else {
		acl.AuthorizedParty = strings.Split(args[0].(string), ",")

		if len(args) > 1 {
			//attributes
			condtions := make(map[string]string)
			attributes := strings.Split(args[1].(string), ",")
			for _, attr := range attributes {
				kv := strings.Split(attr, "=")
				condtions[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
			}
			acl.Conditions = condtions
		}
	}

	handlerinfo.outputs = outputattrs
	handlerinfo.acl = acl

	return &handlerinfo
}

func isPrimitive(dataType string) bool {
	switch dataType {
	case "String":
	case "Long":
	case "Integer":
	case "Double":
	case "DateTime":
	case "Boolean":
		return true
	}

	return false
}
