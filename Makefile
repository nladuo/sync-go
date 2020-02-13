default: build

prepare:
	go get -u golang.org/x/crypto/ssh
	go get -u github.com/pkg/sftp
	go get -u github.com/radovskyb/watcher/...

build:
	go build -o sync-go
