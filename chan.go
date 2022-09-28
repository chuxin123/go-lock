package lock

var locker = make(chan struct{}, 1)

type ChanLock struct {
}

func NewChanLock() *ChanLock {
	locker <- struct{}{}
	return &ChanLock{}
}

func (c *ChanLock) Lock() bool {
	select {
	case <-locker:
		return true
	}
}

func (c *ChanLock) Unlock() bool {
	locker <- struct{}{}
	return true
}
