package main

import (
	"fmt"
	"os"

	"cliapp/cli"
)

func main() {

	app := cli.App()
	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error running app:", err)
	}

}
