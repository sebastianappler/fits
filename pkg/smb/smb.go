package smb

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/hirochachacha/go-smb2"
)

func Send(filename string, data []byte, url url.URL, username string, password string) error {
	port := url.Port()
	if port == "" {
		port = "445"
	}

	conn, err := net.Dial("tcp", url.Hostname()+":"+port)
	if err != nil {
		return fmt.Errorf("unable to connect with smb %v\n", err)
	}
	defer conn.Close()

	// To support anonymous login
	if username == "" && password == "" {
		username = "Guest"
	}

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     username,
			Password: password,
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		return fmt.Errorf("unable to authenticate smb %v\n", err)
	}
	defer s.Logoff()

	pathArr := strings.Split(url.Path, "/")
	share := pathArr[1]
	fs, err := s.Mount(share)
	if err != nil {
		return fmt.Errorf("unable to mount smb share %v\n", err)
	}
	defer fs.Umount()

	sharePath := strings.Join(pathArr[2:], "/")
	if sharePath == "/" {
		sharePath = ""
	}
	f, err := fs.Create(sharePath + filename)
	if err != nil {
		return fmt.Errorf("unable to create file %v\n", err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("unable to write file on smb share %v\n", err)
	}

	fmt.Printf("File sent to smb share %v\n", filename)
	return nil
}
