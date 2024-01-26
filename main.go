package main

import (
	"fmt"
	"go-cli/app"
	"log"
	"os"
)

func main() {
	fmt.Println()
	app := app.Generate()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
