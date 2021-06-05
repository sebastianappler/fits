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
	var fromPathRaw, fromUsername, fromPassword, toPathRaw, toUsername, toPassword string
	config, loadConfigFileErr := toml.LoadFile("./config/config.toml")
	fitsEnvironment := os.Getenv("FITS_ENVIRONMENT")
	fmt.Printf("environment: %v\n", fitsEnvironment)
	// Let config file have precedence over env variables
	if loadConfigFileErr == nil {
		fromPathRaw = config.Get("from.path").(string)
		fromUsername = config.GetDefault("from.username", "").(string)
		fromPassword = config.GetDefault("from.password", "").(string)
		toPathRaw = config.Get("to.path").(string)
		toUsername = config.GetDefault("to.username", "").(string)
		toPassword = config.GetDefault("to.password", "").(string)
	} else {
		fromPathRaw = os.Getenv("FITS_FROM_PATH")
		fromUsername = os.Getenv("FITS_FROM_USERNAME")
		fromPassword = os.Getenv("FITS_FROM_PASSWORD")
		toPathRaw = os.Getenv("FITS_TO_PATH")
		toUsername = os.Getenv("FITS_TO_USERNAME")
		toPassword = os.Getenv("FITS_TO_PASSWORD")

		if(fromPathRaw == "" || toPathRaw == "") {
			fmt.Printf("error loading config file: %v\n", loadConfigFileErr)
			log.Fatal("no config file or environment variables found. Please verify that you have a config.toml file or environment variables FITS_FROM_PATH and FITS_TO_PATH set correctly.")
		}
	}

	// If docker env we need to override the fs paths for container
	if(fitsEnvironment == "docker") {
		if(GetSchemeByUrl(fromPathRaw) == "fs") {
			fromPathRaw = "/from"
		}
		if(GetSchemeByUrl(toPathRaw) == "fs") {
			toPathRaw = "/to"
		}
	}

	fromPath := getPathObj(fromPathRaw, fromUsername, fromPassword)
	fmt.Println("from info:")
	printPath(fromPath)

	toPath := getPathObj(toPathRaw, toUsername, toPassword)
	fmt.Println("to info:")
	printPath(toPath)

	return fromPath, toPath
}

func getPathObj(fromRaw, username, password string) (common.Path) {
	urlRaw := getFullPath(fromRaw)
	urlRaw = strings.Replace(urlRaw, "\\", "/", -1)
	urlParsed, err := url.Parse(urlRaw)
	if err != nil {
		log.Fatal(err)
	}

	path := common.Path{
		Url:      *urlParsed,
		UrlRaw:   urlRaw,
		Username: username,
		Password: password,
	}

	return path
}

func GetSchemeByUrl(urlRaw string) string {
	urlParsed, err := url.Parse(urlRaw)
	if (err != nil) {
		log.Fatal(err)
	}
	if(urlParsed.Scheme != "") {
		return urlParsed.Scheme
	}

	url := strings.Replace(urlRaw, "\\", "/", -1)
	if strings.HasPrefix(url, "//") {
		return "smb"
	}
	return "fs"
}

func getFullPath(path string) string {
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
		scheme = GetSchemeByUrl(path.UrlRaw)
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
