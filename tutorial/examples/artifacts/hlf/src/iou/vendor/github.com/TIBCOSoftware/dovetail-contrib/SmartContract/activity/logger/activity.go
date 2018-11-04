package logger

import (
	impl "github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/activity/logger"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return impl.NewActivity(metadata)
}
