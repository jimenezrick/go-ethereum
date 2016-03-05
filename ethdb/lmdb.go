package ethdb

import (
	"github.com/bmatsuo/lmdb-go/lmdb"
)

type LMDBDatabase struct {
	env *lmdb.Env
}

func NewLMDBDatabase() (*LMDBDatabase, error) {
	env, err := lmdb.NewEnv()
	if err != nil {
		return nil, err
	}

	return &LMDBDatabase{env}, nil
}

func (self *LMDBDatabase) Close() {
	err := self.env.Close()
	if err != nil {
		panic(err)
	}
}
