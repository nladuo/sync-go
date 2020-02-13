package main

import (
	"fmt"
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

func checkFileWriting(path string) bool {

}

// func WatcherTest() {
func main() {
	w := watcher.New()
	// w.SetMaxEvents(1)
	// w.FilterOps(watcher.Create, watcher.Write, watcher.Remove, watcher.Move, watcher.)

	go func() {
		// syncronizing_list := [] 文件列表
		for {
			select {
			case event := <-w.Event:
				// fmt.Println(event) // Print the event's info.
				if !event.IsDir() {
					// 如果再文件列表中已经有了，那么就略过
					//如果没有就先加到列表里，然后再等待删除
					fmt.Println(event.Op, event.Path)

				}

				// fmt.Println(event.Op)
				// fmt.Println(event.OldPath)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive("./test_folder"); err != nil {
		log.Fatalln(err)
	}

	// being watched and their paths.
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	fmt.Println()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
