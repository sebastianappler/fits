package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/sebastianappler/fits/common"
	"github.com/sebastianappler/fits/watcher"
)

func main() {
	config, err := toml.LoadFile("./config/config.toml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config loaded successfully.")

	fromUrlRaw := GetFullPath(config.Get("from.path").(string))
	toUrlRaw := GetFullPath(config.Get("to.path").(string))
	fitsEnvironment := os.Getenv("FITS_ENVIRONMENT")
	fromUrl, err := url.Parse(fromUrlRaw)
	if err != nil {
		log.Fatal(err)
	}
	toUrl, err := url.Parse(toUrlRaw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ENVIRONMENT: %#v\n", fitsEnvironment)
	if fitsEnvironment == "docker" {
		fromUrlRaw = "/from"

		if toUrl.Scheme == "" {
			toUrlRaw = "/to"
		}

		if fromUrl.Scheme == "" {
			fmt.Printf("setting docker from path: %v\n", fromUrlRaw)
			fromUrl, _ = url.Parse(fromUrlRaw)
		}

		if toUrl.Scheme == "" {
			fmt.Printf("setting docker to path: %v\n", toUrlRaw)
			toUrl, _ = url.Parse(toUrlRaw)
		}
	}

	fromPath := common.Path{
		Url:      *fromUrl,
		UrlRaw:   fromUrlRaw,
		Username: config.GetDefault("from.username", "").(string),
		Password: config.GetDefault("from.password", "").(string),
	}
	fmt.Println("FromPath:")
	PrintPath(fromPath)

	toPath := common.Path{
		Url:      *toUrl,
		UrlRaw:   toUrlRaw,
		Username: config.GetDefault("to.username", "").(string),
		Password: config.GetDefault("to.password", "").(string),
	}
	fmt.Println("ToPath:")
	PrintPath(toPath)

	err = watcher.Watch(fromPath, toPath)

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

func PrintPath(path common.Path) {

	scheme := path.Url.Scheme
	user := path.Username
	if user == "" {
		user = path.Url.User.Username()
	}
	host := path.Url.Host
	port := path.Url.Port()
	urlPath := path.Url.Path

	if scheme == "" {
		scheme = "fs"
	}

	fmt.Printf("scheme:%v\n", scheme)

	if user != "" {
		fmt.Printf("user:%v\n", user)
	}

	if host != "" {
		fmt.Printf("host:%v\n", host)
	}

	if port != "" {
		fmt.Printf("port:%v\n", port)
	}

	fmt.Printf("path:%v\n", urlPath)
}
