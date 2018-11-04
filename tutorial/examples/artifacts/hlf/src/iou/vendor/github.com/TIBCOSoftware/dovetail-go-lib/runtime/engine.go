package runtime

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/definition"
	"github.com/TIBCOSoftware/flogo-lib/app"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/engine/runner"
	"github.com/TIBCOSoftware/flogo-lib/util/managed"
)

type DovetailEngine struct {
	app          *app.Config
	actionRunner action.Runner

	trigger trigger.Trigger
}

// New creates a new Engine
func NewEngine(appCfg *app.Config) (*DovetailEngine, error) {
	// App is required
	if appCfg == nil {
		return nil, errors.New("no App configuration provided")
	}
	// Name is required
	if len(appCfg.Name) == 0 {
		return nil, errors.New("no App name provided")
	}
	// Version is required
	if len(appCfg.Version) == 0 {
		return nil, errors.New("no App version provided")
	}

	//fix up app configuration if it is older
	app.FixUpApp(appCfg)

	//add ExplicitReply due to limitation of UI export
	for _, resource := range appCfg.Resources {
		var defRep *definition.DefinitionRep
		err := json.Unmarshal(resource.Data, &defRep)
		if err != nil {
			return nil, err
		}

		defRep.ExplicitReply = true
		data, err := json.Marshal(defRep)
		if err != nil {
			return nil, err
		}

		resource.Data = json.RawMessage(data)
	}

	return &DovetailEngine{app: appCfg}, nil
}

func (e *DovetailEngine) Init() error {

	e.actionRunner = runner.NewDirect()

	actionFactories := action.Factories()
	action.GetFactory(flow.FLOW_REF)

	for _, factory := range actionFactories {
		if initializable, ok := factory.(managed.Initializable); ok {

			if err := initializable.Init(); err != nil {
				return err
			}
		}
	}

	err := app.RegisterResources(e.app.Resources)
	if err != nil {
		return err
	}

	if len(e.app.Triggers) == 0 {
		return fmt.Errorf("There is no trigger defined in the application")
	}

	if len(e.app.Triggers) > 1 {
		return fmt.Errorf("Each application can have only one type of trigger")
	}

	triggers, err := app.CreateTriggers(e.app.Triggers, e.actionRunner)
	if err != nil {
		return fmt.Errorf("Error Creating trigger instances - %s", err.Error())
	}

	for _, t := range triggers {
		e.trigger = t
	}

	return nil
}

func (e *DovetailEngine) GetTrigger() trigger.Trigger {
	return e.trigger
}
