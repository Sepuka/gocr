package def

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gocr/internal/cfg"
)

type (
	containerFnc func(builder *di.Builder, cfg *cfg.Config) error
)

var (
	defs      []containerFnc
	container di.Container
)

func Container() di.Container {
	return container
}

func Register(fnc containerFnc) {
	defs = append(defs, fnc)
}

func Build(cfgPath string) error {
	var (
		builder *di.Builder
		config  *cfg.Config
		fnc     containerFnc
		err     error
	)

	builder, err = di.NewBuilder(di.App, di.Request)
	if err != nil {
		return err
	}

	config, err = cfg.GetConfig(cfgPath)
	if err != nil {
		return err
	}

	for _, fnc = range defs {
		if err = fnc(builder, config); err != nil {
			return err
		}
	}

	container = builder.Build()

	return nil
}
