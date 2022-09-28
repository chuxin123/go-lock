package lock

type Factory interface {
	CreateLock(handler string) Locker
}

var (
	client = new(clientLocker)
)

type Locker interface {
	Lock() bool
	Unlock() bool
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

func SetLock(l Locker) {
	client.store(l)
}

func GetLock() Locker {
	return client.load()
}

func Lock() bool {
	return GetLock().Lock()
}

func Unlock() bool {
	return GetLock().Unlock()
}
