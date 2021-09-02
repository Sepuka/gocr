package ocr

import (
	"github.com/otiai10/gosseract"
	"github.com/sarulabs/di"
	"github.com/sepuka/gocr/def"
	log2 "github.com/sepuka/gocr/def/log"
	"github.com/sepuka/gocr/internal/cfg"
	"github.com/sepuka/gocr/internal/ocr"
	"go.uber.org/zap"
)

const (
	Parser = `def.ocr.parser`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *cfg.Config) error {
		return builder.Add(di.Def{
			Build: func(container di.Container) (interface{}, error) {
				var (
					parser *ocr.Parser
					client = container.Get(Client).(*gosseract.Client)
					log    = container.Get(log2.LoggerDef).(*zap.SugaredLogger)
				)

				parser = ocr.NewParser(client, log)

				return parser, nil
			},
			Name:  Parser,
			Scope: di.App,
			Close: func(obj interface{}) error {
				return obj.(*ocr.Parser).Close()
			},
		})
	})
}
