package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func SftpCopyFile(localpath, remotepath string, config Config) {
	sftpClient, err := connect(config.Username, config.Password, config.Host, config.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()
	remotedir := path.Dir(remotepath)
	// create remote dir
	sftpClient.MkdirAll(remotedir)

	// create destination file
	dstFile, err := sftpClient.Create(remotepath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := os.Open(localpath)
	if err != nil {
		log.Fatal(err)
	}

	// copy source file to destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
}

func SftpDeleteFile(remotepath string, config Config) {
	sftpClient, err := connect(config.Username, config.Password, config.Host, config.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	// delete file
	sftpClient.Remove(remotepath)
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}
