package ethdb

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestNewLMDBDatabase(t *testing.T) {
	file := filepath.Join("/", "tmp", "lmdbtesttmpfile")
	if common.FileExist(file) {
		os.RemoveAll(file)
	}
	db, err := NewLMDBDatabase()
	if err != nil {
		panic(err)
	}

	db.Close()
}
