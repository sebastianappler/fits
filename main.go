package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"github.com/sebastianappler/fits/watchers"
	"github.com/fsnotify/fsnotify"
	"github.com/pelletier/go-toml"
)

func main() {
	config, err := toml.LoadFile("./config.toml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config loaded successfully.")

	fromPath := ""
	toPath := ""
	fitsEnvironment := os.Getenv("FITS_ENVIRONMENT")

	fmt.Printf("ENVIRONMENT: %#v\n", fitsEnvironment)
	if fitsEnvironment == "docker" {
		fromPath = "/from"
		toPath = "/to"
	} else {
		fromPath = GetFullPath(config.Get("from.path").(string))
		toPath = GetFullPath(config.Get("to.path").(string))
	}

	fmt.Printf("Transfering from %#v ", fromPath)
	fmt.Printf("to %#v.\n", toPath)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("Event", event)

				if event.Op&fsnotify.Create == fsnotify.Create {
					moveFrom := event.Name
					moveTo := path.Join(toPath, filepath.Base(event.Name))
					err := watchers.MoveFile(moveFrom, moveTo)
					if err != nil {
						fmt.Println(err)
					}
				}

			case err := <-watcher.Errors:
				log.Fatal(err)
			}
		}
	}()

	if err := watcher.Add(fromPath); err != nil {
		log.Fatal(err)
	}
	<-done
}

func GetFullPath(path string) string {
	strArr := strings.Split(path, "/")
	fullPath := ""
	for _, str := range strArr {
		if str == "" {
			fullPath = fullPath + "/"
		} else if str[0] == '$' {
			fullPath = filepath.Join(fullPath, os.Getenv(str[1:]))
		} else {
			fullPath = filepath.Join(fullPath, str)
		}
	}

	return fullPath
}
