package bolt_backend

import (
	"encoding/json"
	"errors"
	"github.com/P4elme6ka/MyUniRepo"
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"reflect"
	"sort"
)

func (b BoltBackend[T]) SetRecord(rec MyUniRepo.Record[T]) error {
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(value.GetBucket())
		if err != nil {
			return err
		}

		encoded, err := json.Marshal(value)
		if err != nil {
			return err
		}

		id, err := value.GetId().MarshalBinary()
		if err != nil {
			return err
		}

		err = bucket.Put(id, encoded)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (b BoltBackend[T]) GetRecord(query MyUniRepo.Query) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (b BoltBackend[T]) DeleteRecord(query MyUniRepo.Query) error {
	//TODO implement me
	panic("implement me")
}

func (b BoltBackend[T]) UpdateAllRecord(query MyUniRepo.Query, rec MyUniRepo.Record[T]) error {
	//TODO implement me
	panic("implement me")
}

func (b BoltBackend[T]) UpdateFirstRecord(query MyUniRepo.Query, rec MyUniRepo.Record[T]) error {
	//TODO implement me
	panic("implement me")
}

func (b BoltBackend[T]) QueryRecord(query MyUniRepo.Query) ([]*T, error) {
	//TODO implement me
	panic("implement me")
}

func GenericGet[T repository.IModel](db *bolt.DB, uuid uuid.UUID) (*T, error) {
	var res *T = nil
	err := db.View(func(tx *bolt.Tx) error {
		var b T
		bucket := tx.Bucket(b.GetBucket())
		if bucket == nil {
			return errors.New("bucket not found")
		}

		id, err := uuid.MarshalBinary()
		if err != nil {
			return err
		}

		val := bucket.Get(id)

		if val == nil {
			return errors.New("not found")
		}

		err = json.Unmarshal(val, &res)
		if err != nil {
			return err
		}

		return nil
	})
	return res, err
}

func GenericList[T repository.IModel](db *bolt.DB) ([]*T, error) {
	var res []*T = nil
	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		var b T
		bucket := tx.Bucket(b.GetBucket())

		cursor := bucket.Cursor()

		for id, val := cursor.First(); id != nil; id, val = cursor.Next() {
			var serv *T
			err := json.Unmarshal(val, &serv)
			if err != nil {
				return err
			}
			res = append(res, serv)
		}

		return nil
	})
	return res, err
}

func GenericQuery[T repository.IModel](filt repository_query.IFilterRepository, sor repository_query.ISorterRepository, db *bolt.DB, q repository_query.IRepositoryQuery) ([]*T, error) {
	var res []*T = nil
	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		var b T
		bucket := tx.Bucket(b.GetBucket())

		cursor := bucket.Cursor()

		for id, val := cursor.First(); id != nil; id, val = cursor.Next() {
			var serv *T
			err := json.Unmarshal(val, &serv)
			if err != nil {
				return err
			}
			add := true
			for _, filter := range q.GetFilters() {
				if !filt.GetAvailableFilters()[filter.GetOperation()].Apply(repository_query.GetModelFieldByJsonName[T](serv, filter.GetFieldName()), reflect.ValueOf(filter.GetValue())) {
					add = false
				}
			}

			if add {
				res = append(res, serv)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(q.GetSorters()) > 0 {
		sort.Slice(res, func(i, j int) bool {
			for _, s := range q.GetSorters() {
				sortRes := sor.GetAvailableSorters()[s.GetOperation()].Apply(repository_query.GetModelFieldByJsonName[T](res[i], s.GetFieldName()), repository_query.GetModelFieldByJsonName[T](res[j], s.GetFieldName()))
				if repository_query.Equal == sortRes {
					continue
				}
				return repository_query.Equal < sortRes
			}
			return true
		})
	}

	if len(res) > q.GetPagination().GetSkip()+q.GetPagination().GetCount() {
		res = res[q.GetPagination().GetSkip() : q.GetPagination().GetSkip()+q.GetPagination().GetCount()]
	} else if len(res) > q.GetPagination().GetSkip() {
		res = res[q.GetPagination().GetSkip():]
	} else {
		res = nil
	}

	return res, nil
}

func GenericDelete[T repository.IModel](db *bolt.DB, uuid uuid.UUID) error {
	err := db.Update(func(tx *bolt.Tx) error {
		var b T
		bucket := tx.Bucket(b.GetBucket())
		if bucket == nil {
			return errors.New("bucket not found")
		}

		id, err := uuid.MarshalBinary()
		if err != nil {
			return err
		}

		err = bucket.Delete(id)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
