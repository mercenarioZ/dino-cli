package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type spinner struct {
	done      chan struct{}
	wg        sync.WaitGroup
	startedAt time.Time
}

func startSpinner(messages ...string) *spinner {
	if len(messages) == 0 {
		messages = []string{"working..."}
	}

	s := &spinner{done: make(chan struct{}), startedAt: time.Now()}
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
				elapsed := time.Since(s.startedAt).Round(1000 * time.Millisecond)

				messageIndex := int(elapsed/(3*time.Second)) % len(messages)
				message := messages[messageIndex]

				fmt.Fprintf(os.Stderr, "\r\033[2K%s %s [%s]", frames[index], message, elapsed)
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
