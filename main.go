package main

import (
	"fmt"
	"os"
	"time"

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
				Name:    "file",        // Command name
				Aliases: []string{"f"}, // Alias for the command
				Usage:   "Process text files",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "files", // Change the flag name to avoid conflict
						Aliases: []string{"f"},
						Usage:   "File input paths",
					},
					&cli.BoolFlag{
						Name:    "count",
						Aliases: []string{"c"},
						Usage:   "Word count",
					},
					&cli.BoolFlag{
						Name:    "cchar",
						Aliases: []string{"cch"},
						Usage:   "Character count",
					},
					&cli.BoolFlag{
						Name:    "cfreq",
						Aliases: []string{"cf"},
						Usage:   "Word frequency count",
					},
					&cli.BoolFlag{
						Name:    "grep",
						Aliases: []string{"gp"},
						Usage:   "Find common characters",
					},
				},
				Action: func(c *cli.Context) error {
					filePaths := c.StringSlice("files") // Update to use the new flag name
					if len(filePaths) == 0 {
						fmt.Println("Please provide at least one file path using the -files flag.")
						return nil
					}

					if c.Bool("count") {
						wordCount, err := textutil.CountWords(filePaths)
						if err != nil {
							fmt.Println("Error counting words:", err)
							return err
						}
						fmt.Println(wordCount)
					} else if c.Bool("cchar") {
						charCount, err := textutil.CountChar(filePaths)
						if err != nil {
							fmt.Println("Error counting characters:", err)
							return err
						}
						fmt.Println(charCount)
					} else if c.Bool("cfreq") {
						countFreq, err := textutil.CountFreq(filePaths)
						if err != nil {
							fmt.Println("Error counting word frequency:", err)
							return err
						}
						for key, val := range countFreq {
							fmt.Printf("%v: %d\n", key, val)
						}
					} else if c.Bool("grep") {

						start := time.Now()
						res, err := textutil.FindMatchingWords(filePaths)
						elapsed := time.Since(start)

						if err != nil {
							fmt.Println("Error finding matching words:", err)
							return err
						}
						fmt.Println(res)
						fmt.Printf("Execution time : %s", elapsed)
					} else {
						fmt.Println("No valid flags provided. Use -count, -cchar, -cfreq, or -grep.")
						fmt.Println("Flags:", c.FlagNames())
						fmt.Println("Files:", filePaths)

					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error running app:", err)
	}
	// c := make(chan string)
	// go count("sheep", c)
	// msg := <-c
	// fmt.Print(msg)

}

// func count(thing string, c chan string) {
// 	for i := 0; i <= 4; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }
