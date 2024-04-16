package MyUniRepo

type Record[T any] struct {
	Attrs Attributes
	Value T
}
