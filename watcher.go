package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/radovskyb/watcher"
)

func getFileSize(name string) int64 {
	fi, err := os.Stat(name)
	if err != nil {
		return 0
	}
	return fi.Size()
}

func getFileMTime(name string) int64 {
	fi, err := os.Stat(name)
	if err != nil {
		return 0
	}
	mtime := fi.ModTime().UnixNano()
	return mtime
}

func checkFileWriting(path string) (bool, int64) {
	mtime1 := getFileMTime(path)
	time.Sleep(1 * time.Second)
	mtime2 := getFileMTime(path)
	// fmt.Println(mtime1, mtime2)
	return mtime1 != mtime2, mtime2
}

// func WatcherTest() {
func RunWatcher(config Config) {

	// fmt.Println(config)

	w := watcher.New()

	event_to_handle_chan := make(chan watcher.Event, 100)
	event_finish_signal_chan := make(chan string, 100)

	current_dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	color.New(color.FgLightWhite, color.Bold).Print("Watching Dir:")
	color.Green.Println(current_dir)
	color.Cyan.Println("-------------------------")

	// receive files events
	go func() {
		syncronizing_list := []string{} // 文件列表
		for {
			select {
			case event := <-w.Event:
				// filter file
				if strings.HasPrefix(event.Name(), ".") {
					break
				}

				// filter event
				if event.Op == watcher.Chmod {
					break
				}

				if !event.IsDir() {
					not_in := true
					// check if the file is syncronizing
					for _, v := range syncronizing_list {
						if v == event.Path {
							not_in = false
						}
					}
					if not_in {
						syncronizing_list = append(syncronizing_list, event.Path)
						event_to_handle_chan <- event
					}
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case path := <-event_finish_signal_chan:
				index := -1
				//get index
				for i, v := range syncronizing_list {
					if v == path {
						index = i
					}
				}
				// delete from slice
				if index != -1 {
					syncronizing_list = append(syncronizing_list[:index], syncronizing_list[index+1:]...)
				}
			case <-w.Closed:
				return
			}
		}
	}()

	//syncronize the file
	go func() {
		for {
			event := <-event_to_handle_chan
			relative_path := strings.ReplaceAll(event.Path, current_dir, "")
			relative_old_path := strings.ReplaceAll(event.OldPath, current_dir, "")

			time.Sleep(100 * time.Millisecond)

			size := getFileSize(event.Path)
			if size > 1024*20 {
				for {
					is_writing, _ := checkFileWriting(event.Path)
					if !is_writing {
						break
					}
					time.Sleep(100 * time.Millisecond)
				}
			}
			local_filepath := event.Path
			remote_filepath := path.Join(config.RemoteDir, relative_path)
			if (event.Op == watcher.Move) || (event.Op == watcher.Rename) {
				SftpCopyFile(local_filepath, remote_filepath, config)
				fmt.Print("⫸ ")
				color.New(color.FgLightGreen, color.Bold).Print(" Changed ")
				color.Gray.Println(relative_path)

				old_remote_filepath := path.Join(config.RemoteDir, relative_old_path)
				SftpDeleteFile(old_remote_filepath, config)
				fmt.Print("⫸ ")
				color.New(color.FgLightRed, color.Bold).Print(" Deleted ")
				color.Gray.Println(relative_old_path)
			} else if (event.Op == watcher.Create) || (event.Op == watcher.Write) {
				SftpCopyFile(local_filepath, remote_filepath, config)
				fmt.Print("⫸ ")
				color.New(color.FgLightGreen, color.Bold).Print(" Changed ")
				color.Gray.Println(relative_path)

			} else if event.Op == watcher.Remove {
				SftpDeleteFile(remote_filepath, config)
				fmt.Print("⫸ ")
				color.New(color.FgLightRed, color.Bold).Print(" Deleted ")
				color.Gray.Println(relative_path)
			}

			event_finish_signal_chan <- event.Path
		}
	}()

	if err := w.AddRecursive("./"); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
