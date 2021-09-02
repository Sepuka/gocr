package ocr

import (
	"github.com/otiai10/gosseract"
	"go.uber.org/zap"
	"os"
)

type Parser struct {
	client *gosseract.Client
	log    *zap.SugaredLogger
}

func NewParser(
	client *gosseract.Client,
	log *zap.SugaredLogger,
) *Parser {
	return &Parser{
		client: client,
		log:    log,
	}
}

func (p *Parser) Parse(imgPath string) error {
	fl, err := os.Create(`/tmp/gocr.txt`)
	if err != nil {
		return err
	}
	defer fl.Close()

	if err = p.client.SetImage(imgPath); err != nil {
		return err
	}

	if err = p.client.SetTessdataPrefix(`/usr/local/share/tessdata/`); err != nil {
		return err
	}

	output, err := p.client.Text()

	if err != nil {
		p.log.Error(`could not detect text %s`, err)
	}

	_, err = fl.Write([]byte(output))

	return err
}

func (p *Parser) Close() error {
	_ = p.client.Close()
	_ = p.log.Sync()

	return nil
}
