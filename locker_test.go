package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

var (
	counter int64 = 0
	wg            = sync.WaitGroup{}
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
		t.Log("MutexLock test success")
	}
}

func TestChanLock(t *testing.T) {
	counter = 0
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
	if counter != 1000 {
		t.Fatal("ChanLock invalid")
	} else {
		t.Log("ChanLock test success")
	}
}

func TestRedisLock(t *testing.T) {
	locker := NewRedisLock()
	SetLock(locker)

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	Lock()
	resp := redisClient.Get("redis-lock-key")
	if len(resp.Val()) == 0 {
		t.Fatal("Redislock write error")
	} else {
		t.Logf("Redislock value: %s", resp.Val())
	}
	Unlock()
	resp = redisClient.Get("redis-lock-key")
	if len(resp.Val()) == 0 {
		t.Log("RedisLock test success")
	}
}

func TestEtcdLock(t *testing.T) {

	wg.Add(2)
	go func() {
		l1 := NewEtcdLock()
		locker1 := SetLock(l1)

		if !locker1.lock() {
			fmt.Println("Etcdlock1 fail")
		}
		fmt.Println("Goroutine1 get lock")
		time.Sleep(10 * time.Second)
		locker1.unlock()
		fmt.Println("Goroutine1 release lock")
		defer wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		l2 := NewEtcdLock()
		locker2 := SetLock(l2)

		if !locker2.lock() {
			fmt.Println("Etcdlock2 fail")
		}
		fmt.Println("Goroutine2 get lock")
		time.Sleep(1 * time.Second)
		locker2.unlock()
		fmt.Println("Goroutine2 release lock")
		defer wg.Done()
	}()

	wg.Wait()
	t.Log("Etcdlock Test success")
}

func TestAtomicLock(t *testing.T) {
	counter = 0
	locker := NewAtomicLock()
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
		t.Fatal("AutomicLock invalid")
	} else {
		t.Log("AutomicLock test success")
	}
}
