package watchers

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/sebastianappler/fits/common"
	"github.com/sebastianappler/fits/senders"
)

type FsWatcher struct {
	From string
	To   string
}

func FsWatch(fromPath common.Path, toPath common.Path) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)

	fmt.Printf("Transfering from %#v ", fromPath.UrlRaw)
	fmt.Printf("to %#v.\n", toPath.UrlRaw)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("Event", event)

				if event.Op&fsnotify.Create == fsnotify.Create {
					fileLocalPath := event.Name
					senders.Send(fileLocalPath, toPath)
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
