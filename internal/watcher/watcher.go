package watcher

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sebastianappler/fits/internal/common"
	"github.com/sebastianappler/fits/internal/config"
	"github.com/sebastianappler/fits/internal/fileservice"
)

func Watch(fromPath common.Path, toPath common.Path) error {
	fmt.Printf("Transfering from %#v ", fromPath.UrlRaw)
	fmt.Printf("to %#v.\n", toPath.UrlRaw)

	done := make(chan bool)
	s := gocron.NewScheduler(time.Now().Location())
	s.Every(10).Seconds().Do(processFiles, fromPath, toPath)
	s.SingletonMode()
	s.StartAsync()
	<-done

	return nil
}

func processFiles(fromPath common.Path, toPath common.Path) {
	defer handleFileProcessPanic()
	fromSvc := getFileService(fromPath)
	toSvc := getFileService(toPath)

	filenames, err := fromSvc.List(fromPath)

	if err != nil {
		log.Panic(err)
	}

	for _, filename := range filenames {
		fmt.Printf("processing file %v\n", filename)
		data, err := fromSvc.Read(filename, fromPath)
		if err != nil {
			log.Panic(err)
		} else {
			err = toSvc.Send(filename, data, toPath)

			if err != nil {
				log.Panic(err)
			} else {
				fromSvc.Remove(filename, fromPath)
			}
		}
	}
}

func handleFileProcessPanic() {
	if a := recover(); a != nil {
		fmt.Println("recover from panic while processed file.", a)
	}
}

func getFileService(path common.Path) fileservice.FileService {
	scheme := path.Url.Scheme
	if scheme == "" {
		scheme = config.GetSchemeByUrl(path.UrlRaw)
	}

	if scheme == "fs" {
		return fileservice.FsFileService{}
	}
	if scheme == "ftp" {
		return fileservice.FtpFileService{}
	}
	if scheme == "ssh" {
		return fileservice.SshFileService{}
	}
	if scheme == "smb" {
		return fileservice.SmbFileService{}
	}

	return nil
}
