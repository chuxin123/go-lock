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

func (m *mutexLock) Lock() bool {
	lc.Lock()
	return true
}

func (m *mutexLock) Unlock() bool {
	lc.Unlock()
	return true
}
