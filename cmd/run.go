package cmd

import (
	"bufio"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"fmt"
	"github.com/sepuka/gocr/def"
	"github.com/spf13/cobra"
	"os"
)

var (
	text = &cobra.Command{
		Use: `text`,
		RunE: func(cmd *cobra.Command, args []string) error {
			outputText := args[0]
			fl, err := os.Create(`/tmp/gocr.txt`)
			if err != nil {
				return err
			}

			defer fl.Close()

			writer := bufio.NewWriter(fl)
			instance, err := def.Container().SafeGet(def.OCRClient)
			if err != nil {
				return err
			}

			f, err := os.Open(outputText)
			if err != nil {
				return err
			}
			defer f.Close()

			image, err := vision.NewImageFromReader(f)
			if err != nil {
				return err
			}

			ctx := context.Background()
			annotations, _ := instance.(*vision.ImageAnnotatorClient).DetectTexts(ctx, image, nil, 10)

			if len(annotations) == 0 {
				fmt.Fprintln(writer, "No text found.")
			} else {
				fmt.Fprintln(writer, "Text:")
				for _, annotation := range annotations {
					fmt.Fprintf(writer, "%q\n", annotation.Description)
				}
			}

			return nil
		},
		Args: cobra.MinimumNArgs(1),
	}
)

func init() {
	rootCmd.AddCommand(text)
}
