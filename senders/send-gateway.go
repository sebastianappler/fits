package senders

import (
	"fmt"

	"github.com/sebastianappler/fits/common"
)

func Send(fileLocalPath string, toPath common.Path) error {

	scheme := toPath.Url.Scheme

	if scheme == "" {
		scheme = "fs"
	}

	fmt.Printf("scheme:%v\n", scheme)
	fmt.Printf("user:%v\n", toPath.Url.User)
	fmt.Printf("host:%v\n", toPath.Url.Host)
	if toPath.Url.Port() != "" {
		fmt.Printf("port:%v\n", toPath.Url.Host)
	}

	if scheme == "fs" {
		err := FsSend(fileLocalPath, toPath)
		if err != nil {
			return fmt.Errorf("Unable to send file: %v", err)
		}
	}

	if scheme == "ftp" {
		err := FtpSend(fileLocalPath, toPath)
		if err != nil {
			return fmt.Errorf("Unable to send file: %v", err)
		}
	}
	if scheme == "ssh" {
		err := SshSend(fileLocalPath, toPath)
		if err != nil {
			return fmt.Errorf("Unable to send file: %v", err)
		}
	}

	return nil
}
