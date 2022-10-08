package lock

import (
	"sync/atomic"
)

var (
	aclocker int64 = 0
)

type atomicLock struct {
}

func NewAtomicLock() *atomicLock {
	return &atomicLock{}
}

func (l *atomicLock) lock() bool {

	for !atomic.CompareAndSwapInt64(&aclocker, 0, 1) {
	}
	return true
}

func (l *atomicLock) unlock() bool {
	atomic.StoreInt64(&aclocker, 0)
	return true
}
