package bolt_local

import (
	"errors"
	"github.com/P4elme6ka/my_awesome_vpn/repository/repository_query"
	"github.com/labstack/gommon/log"
	"reflect"
	"strings"
)

var boltSortOperationsMap = map[repository_query.SortOperationName]repository_query.ISortOperation{
	repository_query.Descending: boltDescendingSort{},
	repository_query.Ascending:  boltAscendingSort{},
}

func getBoltSortOperationNames() []repository_query.SortOperationName {
	res := make([]repository_query.SortOperationName, 0)
	for i := range boltSortOperationsMap {
		res = append(res, i)
	}
	return res
}

type boltAscendingSort struct{}

func (boltAscendingSort) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) || repository_query.IsTypeFloat(r) || repository_query.IsTypeString(r) {
		return nil
	}
	return errors.New("type not comparable")
}

func (boltAscendingSort) GetName() repository_query.SortOperationName {
	return repository_query.Ascending
}

func (boltAscendingSort) Apply(a reflect.Value, b reflect.Value) repository_query.OperationResult {
	if repository_query.IsTypeInt(a.Type()) {
		if a.Int() == b.Int() {
			return repository_query.Equal
		} else if a.Int() > b.Int() {
			return repository_query.More
		} else {
			return repository_query.Less
		}
	}
	if repository_query.IsTypeString(a.Type()) {
		return repository_query.OperationResult(strings.Compare(a.String(), b.String()))
	}
	if repository_query.IsTypeFloat(a.Type()) {
		if a.Float() == b.Float() {
			return repository_query.Equal
		} else if a.Float() > b.Float() {
			return repository_query.More
		} else {
			return repository_query.Less
		}
	}
	log.Warn("sort operation miss type")
	return repository_query.Equal
}

type boltDescendingSort struct {
}

func (boltDescendingSort) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) || repository_query.IsTypeFloat(r) || repository_query.IsTypeString(r) {
		return nil
	}
	return errors.New("type not comparable")
}

func (boltDescendingSort) GetName() repository_query.SortOperationName {
	return repository_query.Ascending
}

func (boltDescendingSort) Apply(a reflect.Value, b reflect.Value) repository_query.OperationResult {
	if repository_query.IsTypeInt(a.Type()) {
		if a.Int() == b.Int() {
			return repository_query.Equal
		} else if a.Int() > b.Int() {
			return repository_query.More
		} else {
			return repository_query.Less
		}
	}
	if repository_query.IsTypeString(a.Type()) {
		return repository_query.OperationResult(strings.Compare(b.String(), a.String()))
	}
	if repository_query.IsTypeFloat(a.Type()) {
		if a.Float() == b.Float() {
			return repository_query.Equal
		} else if a.Float() < b.Float() {
			return repository_query.More
		} else {
			return repository_query.Less
		}
	}
	log.Warn("sort operation miss type")
	return repository_query.Equal

}
