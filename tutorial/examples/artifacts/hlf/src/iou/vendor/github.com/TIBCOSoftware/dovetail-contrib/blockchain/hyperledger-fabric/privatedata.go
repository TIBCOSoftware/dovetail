package hyperledgerfabric

/*
import (
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/dovetail/runtime/go//SmartContract/common/utils"
)

//field should be in $tx.path.to.field format
type PrivateDataCollection struct {
	name   string
	fields []string
}

//key is asset name
var privateDataRegistry = make(map[string]map[string]PrivateDataCollection)

func RegisterPrivateData(assetName string, pvdata PrivateDataCollection) {
	colls, found := privateDataRegistry[assetName]
	if !found {
		colls = make(map[string]PrivateDataCollection)
		privateDataRegistry[assetName] = colls
	}

	colls[pvdata.name] = pvdata
}

func RegisterPrivateDatas(assetName string, pvdatas []PrivateDataCollection) {
	for _, pvdata := range pvdatas {
		RegisterPrivateData(assetName, pvdata)
	}
}

func ParsePrivateDataDef(assetName, schema string) error {
	parsedSchema, err := utils.ParseSchemaMetadata(schema)
	if err != nil {
		return err
	}

	decorators := utils.GetDecoratorsByPrefix("PrivateData", parsedSchema)
	if len(decorators) > 0 {
		for _, decorator := range decorators {
			args := decorator["args"].([]interface{})

			if len(args) != 2 {
				return fmt.Errorf("Expect 2 arguments, name and fields, for PrivateData")
			}

			pvdata := PrivateDataCollection{name: args[0].(string), fields: make([]string, 0)}
			fields := strings.Split(args[1].(string), ",")
			for _, f := range fields {
				pvdata.fields = append(pvdata.fields, strings.TrimSpace(f))
			}
			RegisterPrivateData(assetName, pvdata)
		}
	}

	return nil
}

func GetPrivateData(assetName, collectionName string) (PrivateDataCollection, bool) {
	if colls, found := privateDataRegistry[assetName]; found {
		if coll, found := colls[collectionName]; found {
			return coll, true
		}
	}

	return PrivateDataCollection{}, false
}

func GetPrivateDatas(assetName string) (map[string]PrivateDataCollection, bool) {
	colls, ok := privateDataRegistry[assetName]
	return colls, ok
}
*/
