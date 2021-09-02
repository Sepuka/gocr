package ocr

import (
	"github.com/otiai10/gosseract"
	"github.com/sarulabs/di"
	"github.com/sepuka/gocr/def"
	"github.com/sepuka/gocr/internal/cfg"
)

const (
	Client = `def.ocr.client`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *cfg.Config) error {
		return builder.Add(di.Def{
			Build: func(container di.Container) (interface{}, error) {
				var (
					client *gosseract.Client
				)

				client = gosseract.NewClient()

				return client, nil
			},
			Name:  Client,
			Scope: di.App,
			Close: func(obj interface{}) error {
				return obj.(*gosseract.Client).Close()
			},
		})
	})
}
