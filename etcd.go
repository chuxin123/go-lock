package lock

import (
	"context"
	"fmt"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	etcdClient  *clientv3.Client
	etcdLockKey = "/etcd-lock-key"
	s           sync.Once
)

type etcdLock struct {
	etcd *clientv3.Client
	mu   *concurrency.Mutex
}

func NewEtcdLock() *etcdLock {

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	})

	if etcdClient == nil || err != nil {
		fmt.Printf("clientv3.New failed")
	}

	session, err := concurrency.NewSession(etcdClient)
	if err != nil {
		return nil
	}
	mutex := concurrency.NewMutex(session, etcdLockKey)
	return &etcdLock{
		etcd: etcdClient,
		mu:   mutex,
	}
}

func (e *etcdLock) Lock() bool {
	err := e.mu.Lock(context.TODO())
	return err == nil
}

func (e *etcdLock) Unlock() bool {
	err := e.mu.Unlock(context.TODO())
	return err == nil
}
