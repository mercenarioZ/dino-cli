package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type spinner struct {
	done chan struct{}
	wg   sync.WaitGroup
}

func startSpinner(message string) *spinner {
	s := &spinner{done: make(chan struct{})}
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()

		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		ticker := time.NewTicker(time.Millisecond * 80)

		defer ticker.Stop()

		index := 0
		for {
			select {
			case <-ticker.C:
				fmt.Fprintf(os.Stderr, "\r%s %s", frames[index], message)
				index = (index + 1) % len(frames)
			case <-s.done:
				fmt.Fprintf(os.Stderr, "\r\033[2K")
				return
			}
		}
	}()

	return s
}

func (s *spinner) stop() {
	close(s.done)
	s.wg.Wait()
}

func statusf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "dino: "+format+"\n", args...)
}
