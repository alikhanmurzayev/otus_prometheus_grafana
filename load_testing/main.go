package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"time"
)

const (
	baseUrl = "http://arch.homework"
)

var totalRequests int64

func main() {
	runtime.GOMAXPROCS(2)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()

	group := errgroup.Group{}

	group.Go(func() error {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		for {
			select {
			case <-interrupt:
				log.Printf("interrupted CTRL+C")
				cancelFunc()
			case <-ctx.Done():
				return nil
			case <-ticker.C:
				logTotalRequests()
			}
		}
	})

	for i := 1; i <= 4; i++ {
		processID := i
		group.Go(func() error {
			return makeRequests(ctx, processID)
		})
	}

	err := group.Wait()

	if errors.Is(err, context.Canceled) {
		log.Println("context cancelled")
	}
	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("context deadline exceeded")
	}
	logTotalRequests()
}

func logTotalRequests() {
	log.Printf("total requests: %d", atomic.LoadInt64(&totalRequests))
}
