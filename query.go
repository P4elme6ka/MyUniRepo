package MyUniRepo

type Query struct {
	Filters    []Filter
	Sorters    []Sorter
	Pagination Pagination
}

type Pagination struct {
	Count int
	Skip  int
}

type FilterOperationName string

const (
	IntMoreThan FilterOperationName = "int_more_than"
	IntLessThan FilterOperationName = "int_less_than"
	IntEqual    FilterOperationName = "int_equal"
	IntNotEqual FilterOperationName = "int_not_equal"

	StringEqual     FilterOperationName = "string_equal"
	StringNotEqual  FilterOperationName = "string_not_equal"
	StringHasSubstr FilterOperationName = "string_has_substr"

	FloatMoreThan FilterOperationName = "float_more_than"
	FloatLessThan FilterOperationName = "float_less_than"
	FloatEqual    FilterOperationName = "float_equal"
	FloatNotEqual FilterOperationName = "float_not_equal"

	IsNull FilterOperationName = "is_null"

	UuidIsEqual FilterOperationName = "uuid_equal"
	UuidIsNull  FilterOperationName = "uuid_is_null"

	TimeMoreThan FilterOperationName = "time_more_than"
	TimeLessThan FilterOperationName = "time_less_than"
	TimeEqual    FilterOperationName = "time_equal"
)

var FilterOperations = []FilterOperationName{
	IntMoreThan,
	IntLessThan,
	IntEqual,
	IntNotEqual,

	StringEqual,
	StringNotEqual,
	StringHasSubstr,

	FloatMoreThan,
	FloatLessThan,
	FloatEqual,
	FloatNotEqual,

	IsNull,

	UuidIsEqual,
	UuidIsNull,

	TimeMoreThan,
	TimeLessThan,
	TimeEqual,
}

type Filter struct {
	Name           FilterOperationName
	CompareToValue any
	FieldName      string
}

type SortOperationName string

const (
	Ascending  SortOperationName = "ascending"
	Descending SortOperationName = "descending"
)

var DefaultSortOperations = []SortOperationName{
	Ascending,
	Descending,
}

type Sorter struct {
	Name      SortOperationName
	FieldName string
}
