package lock

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	counter = 0
	wg      = sync.WaitGroup{}
)

func TestMutexLock(t *testing.T) {
	locker := NewMutexLock()
	SetLock(locker)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			Lock()
			counter++
			Unlock()
			defer wg.Done()
		}()
	}
	wg.Wait()
	if counter != 1000 {
		t.Fatal("MutexLock invalid")
	} else {
		t.Log("MutexLock Test success")
	}
}

func TestChanLock(t *testing.T) {
	locker := NewChanLock()
	SetLock(locker)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			Lock()
			counter++
			Unlock()
			defer wg.Done()
		}()
	}
	wg.Wait()
	if counter != 2000 {
		t.Fatal("ChanLock invalid")
	} else {
		t.Log("ChanLock Test success")
	}
}

func TestRedisLock(t *testing.T) {
	locker := NewRedisLock()
	SetLock(locker)
	// 待完善
	Lock()
	Unlock()
}

func TestEtcdLock(t *testing.T) {
	var counter int64 = 1
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			l := NewEtcdLock()
			if l.Lock() {
				fmt.Println("Locked success:", atomic.LoadInt64(&counter))
				time.Sleep(3 * time.Second)
				l.Unlock()
				atomic.AddInt64(&counter, 1)
				defer wg.Done()
			} else {
				fmt.Println("Locked fail")
			}
		}()
	}
	wg.Wait()
}
