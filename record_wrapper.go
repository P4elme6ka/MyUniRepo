package myGoRepo

type Record[T any] struct {
	Attrs Attributes
	Value T
}
