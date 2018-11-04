package transaction

import (
	impl "github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/trigger/transaction"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

// NewFactory create a new Trigger factory
// Trigger must define this function
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return impl.NewFactory(md)
}
