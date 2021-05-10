package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/sebastianappler/fits/internal/common"
)

func LoadConfig() (common.Path, common.Path) {
	config, err := toml.LoadFile("./config/config.toml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config loaded successfully.")

	fromUrlRaw := getFullPath(config.Get("from.path").(string))
	toUrlRaw := getFullPath(config.Get("to.path").(string))
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
	printPath(fromPath)

	toPath := common.Path{
		Url:      *toUrl,
		UrlRaw:   toUrlRaw,
		Username: config.GetDefault("to.username", "").(string),
		Password: config.GetDefault("to.password", "").(string),
	}
	fmt.Println("ToPath:")
	printPath(toPath)

	return fromPath, toPath
}

func getFullPath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
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

func printPath(path common.Path) {

	scheme := path.Url.Scheme
	user := path.Username
	if user == "" {
		user = path.Url.User.Username()
	}
	host := path.Url.Host
	port := path.Url.Port()
	urlPath := path.Url.Path

	if scheme == "" {
		if strings.HasPrefix(path.UrlRaw, "//") || strings.HasPrefix(path.UrlRaw, "\\\\") {
			scheme = "smb"
		} else {
			scheme = "fs"
		}
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
