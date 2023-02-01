package main

import (
	"log"
	"os"
	"simple-cli-app/app"
)

func main() {
	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
