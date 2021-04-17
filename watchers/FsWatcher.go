package watchers

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type FsWatcher struct {
	From string
	To string
}

func Watch(fromPath string, toPath string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)

	fmt.Printf("Transfering from %#v ", fromPath)
	fmt.Printf("to %#v.\n", toPath)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("Event", event)

				if event.Op&fsnotify.Create == fsnotify.Create {
					moveFrom := event.Name
					moveTo := path.Join(toPath, filepath.Base(event.Name))
					err := moveFile(moveFrom, moveTo)
					if err != nil {
						fmt.Println(err)
					}
				}

			case err := <-watcher.Errors:
				log.Fatal(err)
			}
		}
	}()

	if err := watcher.Add(fromPath); err != nil {
		log.Fatal(err)
	}
	<-done

	return nil
}

func moveFile(from, to string) error {
	in, err := os.Open(from)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}

	out, err := os.Create(to)
	if err != nil {
		in.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}

	defer out.Close()

	_, err = io.Copy(out, in)
	in.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}

	err = out.Sync()
	if err != nil {
		return fmt.Errorf("Sync error: %s", err)
	}

	si, err := os.Stat(from)
	if err != nil {
		return fmt.Errorf("Stat error: %s", err)
	}

	err = os.Chmod(to, si.Mode())
	if err != nil {
		return fmt.Errorf("Chmod error: %s", err)
	}

	err = os.Remove(from)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}

	return nil
}
