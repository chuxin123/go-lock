package lock

import (
	"sync"
)

var (
	lc sync.Mutex
)

type MutexLock struct {
}

func NewMutexLock() *MutexLock {

	return &MutexLock{}
}

func (m *MutexLock) Lock() bool {
	lc.Lock()
	return true
}

func (m *MutexLock) Unlock() bool {
	lc.Unlock()
	return true
}
