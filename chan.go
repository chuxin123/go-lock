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

func (c *chanLock) lock() bool {
	<-locker
	return true
}

func (c *chanLock) unlock() bool {
	locker <- struct{}{}
	return true
}
