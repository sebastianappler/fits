package ssh

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/sftp"
	"github.com/sebastianappler/fits/common"
	"golang.org/x/crypto/ssh"
	kh "golang.org/x/crypto/ssh/knownhosts"
)

func Send(fileName string, fileData []byte, toPath common.Path) error {
	port := toPath.Url.Port()
	if port == "" {
		port = "22"
	}

	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return fmt.Errorf("unable to read known_hosts: %v\n", err)
	}
	defer file.Close()

	hostKeyCallback, err := kh.New(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return fmt.Errorf("could not create hostkeycallback function: %v\n", err)
	}

	config := ssh.ClientConfig{
		User: toPath.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(toPath.Password),
		},
		HostKeyCallback: hostKeyCallback, // ssh.FixedHostKey(hostKey),
	}

	client, err := ssh.Dial("tcp", toPath.Url.Host+":"+port, &config)

	if err != nil {
		return fmt.Errorf("failed to dial: %v\n", err)
	}

	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(client)
	if err != nil {
		return fmt.Errorf("unable to create sftp client: %v\n", err)
	}
	defer sftp.Close()

	// Open the source file
	srcFile := bytes.NewBuffer(fileData)

	// Create the destination file
	dstFile, err := sftp.Create(path.Join(toPath.Url.Path, fileName))
	if err != nil {
		return fmt.Errorf("unable to create destionation file: %v\n", err)
	}

	// write to file
	if _, err := dstFile.ReadFrom(srcFile); err != nil {
		return fmt.Errorf("unable to write file: %v\n", err)
	}

	fmt.Printf("file sent with ssh: %v\n", fileName)
	return nil
}
