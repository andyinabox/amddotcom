package watcher

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/fsnotify/fsnotify"
)

func (w *Watcher) Watch(ctx context.Context, paths []string) (<-chan string, error) {

	// set up channels for channel group
	changes := make(chan string)

	// create fs watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	// cleanup func for when exiting the goroutine
	cleanup := func() {
		watcher.Close()
		close(changes)
	}

	// Start listening for events.
	go func() {
		for {
			select {

			// new change event
			case event, ok := <-watcher.Events:

				// cleanup if there' an issue with the channel
				if !ok {
					err = errors.New("error recieving watcher event")
					cleanup()
					return
				}

				log.Debug("recieved fsnotify event", "event", event.Name, "op", event.Op)

				// skip CHMOD events because it was leading
				// to extra events being fired (and not really what we want)
				if event.Op != fsnotify.Chmod {
					changes <- event.Name
				}

			// new error
			case watchErr, ok := <-watcher.Errors:

				// cleanup if there's an issue with the channel
				if !ok {
					err = errors.New("error recieving watcher error")
					cleanup()
					return
				}
				// otherwise pass on error
				fmt.Printf("error: %s", watchErr)

			// closing out of context
			case <-ctx.Done():
				cleanup()
				return
			}
		}
	}()

	// add paths to watchere
	for _, p := range paths {
		err = watcher.Add(p)
		if err != nil {
			return nil, err
		}
	}

	return changes, nil
}
