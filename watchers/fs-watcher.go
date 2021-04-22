package watchers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/sebastianappler/fits/common"
	"github.com/sebastianappler/fits/senders"
)

func FsWatch(fromPath common.Path, toPath common.Path) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)

	fmt.Printf("Transfering from %#v ", fromPath.UrlRaw)
	fmt.Printf("to %#v.\n", toPath.UrlRaw)

	sendAllFiles(fromPath.UrlRaw, toPath)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("Event", event)

				if event.Op&fsnotify.Create == fsnotify.Create {
					fileLocalPath := event.Name

					fileInfo, err := os.Stat(fileLocalPath)
					if err != nil {
						log.Fatal(err)
					}

					if fileInfo.IsDir() != true {
						senders.Send(fileLocalPath, toPath)
					}
				}

			case err := <-watcher.Errors:
				log.Fatal(err)
			}
		}
	}()

	if err := watcher.Add(fromPath.UrlRaw); err != nil {
		log.Fatal(err)
	}
	<-done

	return nil
}

func sendAllFiles(fromPath string, toPath common.Path) error {

	files, err := ioutil.ReadDir(fromPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() != true {
			senders.Send(filepath.Join(fromPath, file.Name()), toPath)
		}
	}

	return nil
}
