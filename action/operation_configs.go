package action

import (
	"os"

	"github.com/cnabio/cnab-go/driver"
)

// OperationConfigFunc is a configuration function that can be applied to an
// operation.
type OperationConfigFunc func(op *driver.Operation) error

// OperationConfigs is a set of configuration functions that can be applied as a
// unit to an operation.
type OperationConfigs []OperationConfigFunc

// ApplyConfig safely applies the configuration function to the operation, if
// defined, and stops immediately upon the first error.
func (cfgs OperationConfigs) ApplyConfig(op *driver.Operation) error {
	var err error
	for _, cfg := range cfgs {
		err = cfg(op)
		if err != nil {
			return err
		}
	}

	// If outputs weren't configured, default them
	if op.Out == nil {
		op.Out = os.Stdout
	}
	if op.Err == nil {
		op.Err = os.Stderr
	}

	return nil
}
