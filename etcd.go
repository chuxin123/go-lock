package lock

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"sync"
	"time"
)

var (
	etcdClient  *clientv3.Client
	etcdLockKey = "/etcd-lock-key"
	s           sync.Once
)

type EtcdLock struct {
	etcd *clientv3.Client
	mu   *concurrency.Mutex
}

func NewEtcdLock() *EtcdLock {

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	})

	if etcdClient == nil || err != nil {
		fmt.Printf("etcdsync.New failed")
	}

	session, err := concurrency.NewSession(etcdClient)
	if err != nil {
		return nil
	}
	mutex := concurrency.NewMutex(session, etcdLockKey)
	return &EtcdLock{
		etcd: etcdClient,
		mu:   mutex,
	}
}

func (e *EtcdLock) Lock() bool {
	err := e.mu.Lock(context.TODO())
	if err != nil {
		return false
	}
	return true
}

func (e *EtcdLock) Unlock() bool {
	err := e.mu.Unlock(context.TODO())
	if err != nil {
		return false
	}
	return true
}
