// Code generated from Pkl module `paropd.config.ComputedConfig`. DO NOT EDIT.
package computed

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type ComputedConfig struct {
	Uuid string `pkl:"uuid"`

	HostName string `pkl:"hostName"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a ComputedConfig
func LoadFromPath(ctx context.Context, path string) (ret *ComputedConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a ComputedConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*ComputedConfig, error) {
	var ret ComputedConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
