package main

import (
	"log"

	"github.com/sebastianappler/fits/internal/config"
	"github.com/sebastianappler/fits/internal/watcher"
)

func main() {

	fromPath, toPath := config.LoadConfig()
	err := watcher.Watch(fromPath, toPath)

	if err != nil {
		log.Fatal(err)
	}
}
