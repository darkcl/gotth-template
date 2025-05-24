package store

import (
	"gotth/utils"
	"time"

	bolt "go.etcd.io/bbolt"
)

type BBoltDatabase struct {
	DB *bolt.DB
}

func MustNewBoltDB(path string) *BBoltDatabase {
	db, err := bolt.Open(path, 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})

	utils.CheckError(err)
	return &BBoltDatabase{
		DB: db,
	}
}

func (bb *BBoltDatabase) Close() error {
	return bb.DB.Close()
}
