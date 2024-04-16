package bolt_backend

import (
	"github.com/P4elme6ka/MyUniRepo"
	"log"
)

type BoltBackend struct {
	Path string
}

func NewBoltRepository(conf *MyUniRepo.Settings) *BoltBackend {
	log.Println("Opening bolt store: ", conf.DbPath)
	db, err := bolt.Open(conf.DbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &BoltBackend{}
}
