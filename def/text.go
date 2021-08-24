package def

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"github.com/sarulabs/di"
	"github.com/sepuka/gocr/internal/cfg"
)

const (
	OCRClient = `def.ocr.client`
)

func init() {
	Register(func(builder *di.Builder, cfg *cfg.Config) error {
		return builder.Add(di.Def{
			Build: func(container di.Container) (interface{}, error) {
				var (
					client *vision.ImageAnnotatorClient
				)
				ctx := context.Background()

				client, err := vision.NewImageAnnotatorClient(ctx)
				if err != nil {
					return nil, err
				}

				return client, nil
			},
			Name:  OCRClient,
			Scope: di.App,
			Close: func(obj interface{}) error {
				return obj.(*vision.ImageAnnotatorClient).Close()
			},
		})
	})
}
