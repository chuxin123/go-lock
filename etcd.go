package lock

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	etcdClient  *clientv3.Client
	etcdLockKey = "/etcd-lock-key"
)

type etcdLock struct {
	mu *concurrency.Mutex
}

func init() {
	var err error
	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	})

	if etcdClient == nil || err != nil {
		fmt.Printf("clientv3.New failed")
	}
}

func NewEtcdLock() *etcdLock {

	session, err := concurrency.NewSession(etcdClient)
	if err != nil {
		return nil
	}
	mutex := concurrency.NewMutex(session, etcdLockKey)
	return &etcdLock{
		mu: mutex,
	}
}

func (e *etcdLock) lock() bool {
	err := e.mu.Lock(context.TODO())
	return err == nil
}

func (e *etcdLock) unlock() bool {
	err := e.mu.Unlock(context.TODO())
	return err == nil
}
