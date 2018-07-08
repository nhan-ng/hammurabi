package main

import (
	"log"

	"github.com/nhan-ng/hammurabi/pkg/cmd"
)

func main() {
	if err := cmd.NewHammurabiCmd().Execute(); err != nil {
		log.Println(err)
	}
}
