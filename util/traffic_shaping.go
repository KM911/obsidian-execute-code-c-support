package util

import (
	"sync"
	"time"
)

func NewDebounce(interval time.Duration) func(f func()) {
	var l sync.Mutex
	var timer *time.Timer

	return func(f func()) {
		l.Lock()
		defer l.Unlock()
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(interval, f)
	}
}

func NewThrottle(interval time.Duration) func(f func()) {
	var l sync.Mutex
	var timer *time.Timer

	return func(f func()) {
		l.Lock()
		defer l.Unlock()
		if timer == nil {
			timer = time.AfterFunc(interval, func() {
				f()
				timer = nil
			})
		}
	}
}
