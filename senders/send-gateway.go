package senders

import (
	"fmt"

	"github.com/sebastianappler/fits/common"
)

func Send(fileLocalPath string, toPath common.Path) {

	scheme := toPath.Url.Scheme
	if scheme == "" {
		err := FsSend(fileLocalPath, toPath)
		if err != nil {
			fmt.Println(err)
		}
	}

	if scheme == "ftp" {
		err := FtpSend(fileLocalPath, toPath)
		if err != nil {
			fmt.Println(err)
		}
	}
}
