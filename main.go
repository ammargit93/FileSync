package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func countChar(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	file.Close()

	return len(data), nil
}

func countWords(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	text := strings.TrimSpace(string(data)) // Trim extra spaces
	if text == "" {                         // If file is empty, return count as 0
		return 0, nil
	}

	words := strings.Fields(text) // Fields splits by any whitespace, handling multiple spaces correctly
	return len(words), nil
}

func main() {

	app := &cli.App{
		Name:    "filesync",
		Usage:   "A simple CLI for text processing and matching",
		Version: "1.0.0",

		Commands: []*cli.Command{
			{
				Name:    "count",
				Aliases: []string{"c"},
				Usage:   "Returns Word count from a text file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "Path to the file",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					filepath := c.String("file")
					res, err := countWords(filepath)
					fmt.Println(res)
					return err
				},
			},

			{
				Name:    "cchar",
				Aliases: []string{"cch"},
				Usage:   "Returns char count from a text file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "Path to the file",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					filepath := c.String("file")
					res, err := countChar(filepath)
					fmt.Println(res)
					return err
				},
			},
		},
	}

	// Run the application
	var err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
