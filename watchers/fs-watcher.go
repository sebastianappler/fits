package watchers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sebastianappler/fits/common"
	"github.com/sebastianappler/fits/senders"
)

func FsWatch(fromPath common.Path, toPath common.Path) error {
	fmt.Printf("Transfering from %#v ", fromPath.UrlRaw)
	fmt.Printf("to %#v.\n", toPath.UrlRaw)

	done := make(chan bool)
	s := gocron.NewScheduler(time.Now().Location())
	s.Every(10).Seconds().Do(sendAllFiles, fromPath.UrlRaw, toPath)
	s.StartAsync()
	<-done

	return nil
}

func sendAllFiles(fromPath string, toPath common.Path) error {

	files, err := ioutil.ReadDir(fromPath)
	if err != nil {
		return fmt.Errorf("unable to read files: %v\n", err)
	}

	for _, file := range files {
		if file.IsDir() != true {
			fileLocalPath := filepath.Join(fromPath, file.Name())
			err = senders.Send(fileLocalPath, toPath)

			if err != nil {
				log.Fatalf("Failed to send file: %v", err)
			} else {
				// Remove file when sent
				err = os.Remove(fileLocalPath)
				if err != nil {
					log.Fatalf("Failed removing original file: %v", err)
				}
			}
		}
	}

	return nil
}
