package main

import (
	"log"

	"github.com/nhan-ng/hammurabi/cmd/console-player/app"
)

func main() {
	if err := app.NewHammurabiCmd().Execute(); err != nil {
		log.Println(err)
	}
}
