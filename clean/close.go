package clean

import (
	"sync"
)

var (
	closes  []func() error
	closeMu sync.Mutex
)

func PushClose(close func() error) {
	closeMu.Lock()
	defer closeMu.Unlock()
	closes = append(closes, close)
}

func Close() (errs Errs) {
	closeMu.Lock()
	defer closeMu.Unlock()
	for len(closes) > 0 {
		i := len(closes) - 1
		if err := closes[i](); err != nil {
			errs = append(errs, err)
		}
		closes = closes[:i]
	}
	return
}
