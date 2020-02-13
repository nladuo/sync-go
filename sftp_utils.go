package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func CopyFile(localpath, remotepath string) {

}

func DeleteFile(remotepath string) {

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

func TestSCP() {
	sftpClient, err := connect("ljm", "l82566258", "tengxun", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	// create destination file
	dstFile, err := sftpClient.Create("/home/ljm/main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	err = sftpClient.Mkdir("3232")
	if err != nil {
		// log.Fatal(err)
	}

	// sftpClient.MkdirAll("/home/ljm/dsd/sdsds/232")

	err = sftpClient.RemoveDirectory("/home/ljm/dsd")
	if err != nil {
		fmt.Println(err)
	}

	sftpClient.Remove("/home/ljm/test.txt")

	// create source file
	srcFile, err := os.Open("./main.go")
	if err != nil {
		log.Fatal(err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)
}
