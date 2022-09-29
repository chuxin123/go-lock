package lock

import (
	"sync"
)

var (
	lc sync.Mutex
)

type mutexLock struct {
}

func NewMutexLock() *mutexLock {

	return &mutexLock{}
}

func (m *mutexLock) lock() bool {
	lc.Lock()
	return true
}

func (m *mutexLock) unlock() bool {
	lc.Unlock()
	return true
}
