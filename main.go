package main

import (
	"fmt"
	"os"

	"cliapp/textutil"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:    "filesync",
		Usage:   "A simple CLI for text processing and matching",
		Version: "1.0.0",

		Commands: []*cli.Command{
			{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Returns Word count from a text file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "file",
						Aliases: []string{"f"},
						Usage:   "file input",
					},
					&cli.BoolFlag{
						Name:    "count",
						Aliases: []string{"c"},
						Usage:   "Word Count",
					},
					&cli.BoolFlag{
						Name:    "cchar",
						Aliases: []string{"cch"},
						Usage:   "Char count",
					},
					&cli.BoolFlag{
						Name:    "cfreq",
						Aliases: []string{"cf"},
						Usage:   "Word Count frequency",
					},
				},
				Action: func(c *cli.Context) error {
					filepath := c.String("file")
					if filepath == "" {
						fmt.Println("Please provide a valid file path using the -f flag.")
						return nil
					}
					if c.Bool("count") {
						wordCount, _ := textutil.CountWords(filepath)
						fmt.Println(wordCount)
					} else if c.Bool("cch") {
						charCount, _ := textutil.CountChar(filepath)
						fmt.Println(charCount)
					} else if c.Bool("cfreq") {
						countFreq, _ := textutil.CountFreq(filepath)
						for key, val := range countFreq {
							fmt.Printf("%v: %d\n", key, val)
						}
					}

					return nil
				},
			},
		},
	}

	var err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
