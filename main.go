package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/sebastianappler/fits/common"
	"github.com/sebastianappler/fits/watchers"
)

func main() {
	config, err := toml.LoadFile("./config.toml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config loaded successfully.")

	fromUrlRaw := ""
	toUrlRaw := ""
	fitsEnvironment := os.Getenv("FITS_ENVIRONMENT")

	fmt.Printf("ENVIRONMENT: %#v\n", fitsEnvironment)
	if fitsEnvironment == "docker" {
		fromUrlRaw = "/from"
		toUrlRaw = "/to"
	} else {
		fromUrlRaw = GetFullPath(config.Get("from.path").(string))
		toUrlRaw = GetFullPath(config.Get("to.path").(string))
	}

	fromUrl, err := url.Parse(fromUrlRaw)
	if err != nil {
		log.Fatal(err)
	}
	fromPath := common.Path{
		Url:      *fromUrl,
		UrlRaw:   fromUrlRaw,
		Username: config.GetDefault("from.username", "").(string),
		Password: config.GetDefault("from.password", "").(string),
	}

	toUrl, err := url.Parse(toUrlRaw)
	if err != nil {
		log.Fatal(err)
	}
	toPath := common.Path{
		Url:      *toUrl,
		UrlRaw:   toUrlRaw,
		Username: config.GetDefault("to.username", "").(string),
		Password: config.GetDefault("to.password", "").(string),
	}

	err = watchers.FsWatch(fromPath, toPath)
	if err != nil {
		log.Fatal(err)
	}
}

func GetFullPath(path string) string {
	strArr := strings.Split(path, "/")
	fullPath := ""
	for _, str := range strArr {
		if str == "" {
			fullPath += "/"
		} else if str[0] == '$' {
			fullPath = fullPath + os.Getenv(str[1:]) + "/"
		} else {
			fullPath = fullPath + str + "/"
		}
	}

	return fullPath
}
