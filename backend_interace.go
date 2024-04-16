package MyUniRepo

type IBackendRepo[T any] interface {
	SetRecord(rec Record[T]) error
	GetRecord(query Query) (*T, error)
	DeleteRecord(query Query) error
	UpdateAllRecord(query Query, rec Record[T]) error
	UpdateFirstRecord(query Query, rec Record[T]) error
	QueryRecord(query Query) ([]*T, error)
}
