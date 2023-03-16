package clean

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	cancels  []func()
	cancelMu sync.Mutex
)

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(
		c, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGHUP,
	)
	go func() {
		<-c
		Cancel()
	}()
}

func PushCancel(cancel func()) {
	cancelMu.Lock()
	defer cancelMu.Unlock()
	cancels = append(cancels, cancel)
}

func Cancel() {
	cancelMu.Lock()
	defer cancelMu.Unlock()
	for len(cancels) > 0 {
		cancels[len(cancels)-1]()
		cancels = cancels[:len(cancels)-1]
	}
}
