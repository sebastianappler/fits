package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sebastianappler/fits/watchers"
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

	err = watchers.Watch(fromPath, toPath)
	if err != nil {
		log.Fatal(err)
	}
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

