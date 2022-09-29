package lock

var (
	locker = make(chan struct{}, 1)
)

type chanLock struct {
}

func NewChanLock() *chanLock {
	locker <- struct{}{}
	return &chanLock{}
}

func (c *chanLock) Lock() bool {
	<-locker
	return true
}

func (c *chanLock) Unlock() bool {
	locker <- struct{}{}
	return true
}
