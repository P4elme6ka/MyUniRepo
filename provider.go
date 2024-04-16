package myGoRepo

import "errors"

type Provider string

const (
	BoltBackend  Provider = "bolt_prov"
	MongoBackend Provider = "mongo_prov"
)

func CreateRepo[T any](provider Provider, settings Settings) (*Repo[T], error) {
	switch provider {
	case BoltBackend:
		return nil, errors.New("")
	case MongoBackend:
		return nil, errors.New("")
	default:
		return nil, errors.New("erwerweerqwrqw")
	}
}
