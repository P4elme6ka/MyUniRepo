package MyUniRepo

import "reflect"

type Repo[T any] struct {
	modelType   reflect.Type
	backendRepo IBackendRepo[T]
}

func (rb *Repo[T]) Get(query Query) (*T, error) {
	return nil, nil
}

func (rb *Repo[T]) Set(model T) error {
	return nil
}

func (rb *Repo[T]) Delete(query Query) error {
	return nil
}

func (rb *Repo[T]) Update(query Query, model T) error {
	return nil
}
