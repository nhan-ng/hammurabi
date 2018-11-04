package main

import (
	"log"

	"github.com/nhan-ng/hammurabi/cmd/cli/app"
)

func main() {
	if err := app.NewHammurabiCmd().Execute(); err != nil {
		log.Println(err)
	}
}
