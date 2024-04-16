package bolt_local

import (
	"log"
	"myGoRepo"
)

type BoltBackend struct {
	Path string
}

func NewBoltRepository(conf *myGoRepo.Settings) *BoltBackend {
	log.Println("Opening bolt store: ", conf.DbPath)
	db, err := bolt.Open(conf.DbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &BoltBackend{}
}
