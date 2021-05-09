package watcher

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/internal/fileservice"
)

func Watch(fromPath common.Path, toPath common.Path) error {
	fmt.Printf("Transfering from %#v ", fromPath.UrlRaw)
	fmt.Printf("to %#v.\n", toPath.UrlRaw)

	done := make(chan bool)
	s := gocron.NewScheduler(time.Now().Location())
	s.Every(10).Seconds().Do(ProcessFiles, fromPath, toPath)
	s.StartAsync()
	<-done

	return nil
}

func ProcessFiles(fromPath common.Path, toPath common.Path) {
	fromSvc := GetFileService(fromPath.Url.Scheme)
	toSvc := GetFileService(toPath.Url.Scheme)

	fileNames, err := fromSvc.List(fromPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range fileNames {
		fmt.Printf("processing file %v\n", fileName)
		fileData, err := fromSvc.Read(fileName, fromPath)
		if err != nil {
			log.Fatal(err)
		} else {
			err = toSvc.Send(fileName, fileData, toPath)

			if err != nil {
				log.Fatal(err)
			} else {
				fromSvc.Remove(fileName, fromPath)
			}
		}
	}
}

func GetFileService(scheme string) fileservice.FileService {
	if scheme == "" {
		return fileservice.FsFileService{}
	}
	if scheme == "ftp" {
		return fileservice.FtpFileService{}
	}
	if scheme == "ssh" {
		return fileservice.SshFileService{}
	}

	return nil
}
