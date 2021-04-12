package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/pelletier/go-toml"
)

func main() {

	config, err := toml.LoadFile("./config.toml")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Config loaded successfully")

	fromPath := GetFullPath(config.Get("from.path").(string))
	fmt.Printf("Transfering from %#v to ", fromPath)
	toPath := GetFullPath(config.Get("to.path").(string))
	fmt.Printf("%#v\n", toPath)

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
					err := MoveFile(moveFrom, moveTo)
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

func MoveFile(from, to string) error {
	in, err := os.Open(from)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}

	out, err := os.Create(to)
	if err != nil {
		in.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}

	defer out.Close()

	_, err = io.Copy(out, in)
	in.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}

	err = out.Sync()
	if err != nil {
		return fmt.Errorf("Sync error: %s", err)
	}

	si, err := os.Stat(from)
	if err != nil {
		return fmt.Errorf("Stat error: %s", err)
	}

	err = os.Chmod(to, si.Mode())
	if err != nil {
		return fmt.Errorf("Chmod error: %s", err)
	}

	err = os.Remove(from)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}

	return nil
}
