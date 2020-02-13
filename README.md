# sync-go
sync files between multiple remote servers using sftp.

## Why this Project
For common usage, the sftp syncronizer can be only used in local directory. If you use sftp-tool from a remote server to another remote server, the condition would be complicated. Between the remote server with network latency, I need to syncronize a file after I make sure the file is not being writed.

For example, when I use [https://github.com/serkanyersen/sync](https://github.com/serkanyersen/sync) in local folder, it works well. But my purpose is to sync my local dir to a remote dir, and then sync this remote dir to another remote dir. When doing this, I found most of files would be zero size afther syncronizing. And I am aware of the network latency problem. So I create this project.

## Install Dependencies
``` bash
make prepare
```

## Build Project
``` bash
make
```


## LICENSE
MIT
