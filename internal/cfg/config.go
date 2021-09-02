package cfg

import (
	"github.com/stevenroose/gonfig"
)

type (
	Log struct {
		Prod bool
	}

	Output struct {
		Path string
	}

	Config struct {
		Log Log
		Output Output
	}
)

func GetConfig(path string) (*Config, error) {
	var (
		cfg = &Config{}
		err = gonfig.Load(cfg, gonfig.Conf{
			FileDefaultFilename: path,
			FlagIgnoreUnknown:   true,
			FlagDisable:         true,
			EnvDisable:          true,
		})
	)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
