package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/app"
	"github.com/TIBCOSoftware/flogo-lib/app/resource"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	dt_runtime "github.com/TIBCOSoftware/dovetail-go-lib/runtime"
	dt_trigger "github.com/TIBCOSoftware/dovetail-go-lib/runtime/trigger"
	
	dt_transaction "github.com/TIBCOSoftware/dovetail-contrib/SmartContract/trigger/transaction"
	
	
	dt_ledger "github.com/TIBCOSoftware/dovetail-contrib/SmartContract/activity/ledger"
	
	dt_logger "github.com/TIBCOSoftware/dovetail-contrib/SmartContract/activity/logger"
	
	dt_query "github.com/TIBCOSoftware/dovetail-contrib/SmartContract/activity/query"
	
	dt_txnreply "github.com/TIBCOSoftware/dovetail-contrib/SmartContract/activity/txnreply"
	
)

var enableSecurity = false
func init() {
	fmt.Println("init")
	var cp = EmbeddedProvider()

	appConfig, err := cp.GetApp()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = registerActivities(appConfig.Resources)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = registerTriggers()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	logger.Info("Create DovetailEngine...")
	e, err := dt_runtime.NewEngine(appConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	logger.Info("Init DovetailEngine ...")
	err = e.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	logger.Info("Init DovetailEngine ... Done")

	flowcc.TxnTrigger = e.GetTrigger().(dt_trigger.SmartContractTrigger)
}

func registerTriggers() error {
	logger.Info("registerTrigger...")
	
	
	respath := "vendor/github.com/TIBCOSoftware/dovetail-contrib/SmartContract/trigger/transaction/trigger.json"
	jsonbytes, err := Asset(respath)
	if err != nil {
		return err
	}

	metadata := trigger.Metadata{}
	err = metadata.UnmarshalJSON(jsonbytes)
	if err != nil {
		return err
	}

	trigger.RegisterFactory("github.com/TIBCOSoftware/dovetail-contrib/SmartContract/trigger/transaction", dt_transaction.NewFactory(&metadata))
	
	
	logger.Info("registerTrigger... Done")
	return nil
}

func registerActivities(rConfigs []*resource.Config) error {
	logger.Info("registerActivities...")

	activities := make(map[string]string)
	//template
	
	activities["ledger"] = "github.com/TIBCOSoftware/dovetail-contrib/SmartContract/activity/ledger"
	
	activities["logger"] = "github.com/TIBCOSoftware/dovetail-contrib/SmartContract/activity/logger"
	
	activities["query"] = "github.com/TIBCOSoftware/dovetail-contrib/SmartContract/activity/query"
	
	activities["txnreply"] = "github.com/TIBCOSoftware/dovetail-contrib/SmartContract/activity/txnreply"
	

	acmetadata, err := loadMetadata(activities)
	if err != nil {
		return err
	}

	//template
	
	activity.Register(dt_ledger.NewActivity(acmetadata["ledger"]))
	
	activity.Register(dt_logger.NewActivity(acmetadata["logger"]))
	
	activity.Register(dt_query.NewActivity(acmetadata["query"]))
	
	activity.Register(dt_txnreply.NewActivity(acmetadata["txnreply"]))
	

	logger.Info("registerActivities... Done")
	return nil
}

func loadMetadata(refs map[string]string) (map[string]*activity.Metadata, error) {
	logger.Info("loadMetadata...")
	configs := make(map[string]*activity.Metadata)
	for nm, ref := range refs {
		respath := "vendor/" + ref + "/activity.json"
		jsonbytes, err := Asset(respath)
		if err != nil {
			return nil, err
		}

		meta := &activity.Metadata{}
		err = meta.UnmarshalJSON(jsonbytes)
		if err != nil {
			return nil, err
		}

		configs[nm] = meta
	}
	logger.Info("loadMetadata... Done")
	return configs, nil
}

// embeddedConfigProvider implementation of ConfigProvider
type embeddedProvider struct {
}

//EmbeddedProvider returns an app config from a compiled json file
func EmbeddedProvider() app.ConfigProvider {
	return &embeddedProvider{}
}

// GetApp returns the app configuration
func (d *embeddedProvider) GetApp() (*app.Config, error) {

	appCfg := &app.Config{}

	flowjson, err := Asset("IOU.json")
	if err != nil {
		return nil, err
	}
	jsonParser := json.NewDecoder(bytes.NewReader(flowjson))
	err = jsonParser.Decode(&appCfg)
	if err != nil {
		return nil, err
	}

	return appCfg, nil
}
