package watcher

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/internal/gateway"
)

func Watch(fromPath common.Path, toPath common.Path) error {
	fmt.Printf("Transfering from %#v ", fromPath.UrlRaw)
	fmt.Printf("to %#v.\n", toPath.UrlRaw)

	done := make(chan bool)
	s := gocron.NewScheduler(time.Now().Location())
	s.Every(10).Seconds().Do(SendAllFiles, fromPath, toPath)
	s.StartAsync()
	<-done

	return nil
}

func SendAllFiles(fromPath common.Path, toPath common.Path) {
	fileNames, err := gateway.GetAllFileNames(fromPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range fileNames {
		fmt.Printf("processing file %v\n", fileName)
		fileData, err := gateway.ReadFile(fileName, fromPath)
		if err != nil {
			log.Fatal(err)
		} else {
			err = gateway.Send(fileName, fileData, toPath)

			if err != nil {
				log.Fatal(err)
			} else {
				gateway.Remove(fileName, fromPath)
			}
		}
	}
}
