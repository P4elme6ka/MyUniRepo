package bolt_backend

import (
	"github.com/P4elme6ka/MyUniRepo"
	"github.com/boltdb/bolt"
	"log"
)

type BoltBackend[T any] struct {
	Path string
	db   *bolt.DB
}

func NewBoltRepository[T any](conf *MyUniRepo.Settings) *BoltBackend[T] {
	log.Println("Opening bolt store: ", conf.ConnectionString)
	db, err := bolt.Open(conf.DbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &BoltBackend[T]{
		Path: conf.ConnectionString,
		db:   db,
	}
}
