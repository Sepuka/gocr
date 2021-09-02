package cmd

import (
	"github.com/sepuka/gocr/def"
	"github.com/sepuka/gocr/def/ocr"
	ocr2 "github.com/sepuka/gocr/internal/ocr"
	"github.com/spf13/cobra"
)

var (
	text = &cobra.Command{
		Use:     `text`,
		Aliases: []string{`parse`},
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				err        error
				targetFile = args[0]
			)

			parser, err := def.Container().SafeGet(ocr.Parser)
			if err != nil {
				return err
			}

			return parser.(*ocr2.Parser).Parse(targetFile)
		},
		Args: cobra.MinimumNArgs(1),
	}
)

func init() {
	rootCmd.AddCommand(text)
}
