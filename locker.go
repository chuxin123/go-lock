package lock

var (
	client = new(clientLocker)
)

type Locker interface {
	lock() bool
	unlock() bool
}

type clientLocker struct {
	locker Locker
}

func (c *clientLocker) store(v Locker) {
	c.locker = v
}

func (c *clientLocker) load() Locker {
	return c.locker
}

func SetLock(l Locker) Locker {
	client.store(l)
	return l
}

func GetLock() Locker {
	return client.load()
}

func Lock() bool {
	return GetLock().lock()
}

func Unlock() bool {
	return GetLock().unlock()
}
