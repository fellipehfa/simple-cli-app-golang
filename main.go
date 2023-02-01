package main

import (
	"fmt"
	"log"
	"os"
	"simple-cli-app/app"
)

func main() {
	fmt.Println("Start Point!")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
