package db

import (
	"errors"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	once sync.Once
)

type Instance struct {
	db *leveldb.DB
	mu sync.Mutex
}

func GetInstance() *Instance {
	in := Instance{}
	once.Do(func() {
		var err error
		in.db, err = leveldb.OpenFile("data/allstars.db", nil)
		if err != nil {
			panic(err)
		}
	})
	return &in
}

func (in *Instance) Close() error {
	in.mu.Lock()
	defer in.mu.Unlock()
	if in.db != nil {
		return in.db.Close()
	}
	return nil
}

func (in *Instance) Get(key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("empty key")
	}
	res, err := in.db.Get(key, nil)
	if err == leveldb.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return res, nil
}

func (in *Instance) Set(key, value []byte) error {
	if len(key) == 0 {
		return errors.New("empty key")
	}
	return in.db.Put(key, value, nil)
}
