package bolt_backend

import (
	"errors"
	"github.com/P4elme6ka/my_awesome_vpn/repository/repository_query"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

const (
	UuidEqual repository_query.FilterOperationName = "uuid_equal"
)

var boltFilterOperationsMap = map[repository_query.FilterOperationName]repository_query.IFilterOperation{
	repository_query.IntMoreThan: BoltIntMoreThan{},
	repository_query.IntLessThan: BoltIntLessThan{},
	repository_query.IntEqual:    BoltIntEqual{},
	repository_query.IntNotEqual: BoltIntNotEqual{},

	repository_query.StringEqual:     BoltStringEqual{},
	repository_query.StringNotEqual:  BoltStringNotEqual{},
	repository_query.StringHasSubstr: BoltStringHasSubstr{},

	repository_query.FloatMoreThan: BoltFloatMoreThan{},
	repository_query.FloatLessThan: BoltFloatLessThan{},
	repository_query.FloatEqual:    BoltFloatEqual{},
	repository_query.FloatNotEqual: BoltFloatNotEqual{},

	UuidEqual: BoltUuidEqual{},

	repository_query.IsNull: BoltIsNull{},
}

func getBoltFilterOperationNames() []repository_query.FilterOperationName {
	res := make([]repository_query.FilterOperationName, 0)
	for i := range boltFilterOperationsMap {
		res = append(res, i)
	}
	return res
}

type BoltIntMoreThan struct{}

func (BoltIntMoreThan) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) {
		return nil
	}
	return errors.New("needs an integer")
}

func (BoltIntMoreThan) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) {
		return nil
	}
	return errors.New("needs an integer")
}

func (BoltIntMoreThan) GetName() repository_query.FilterOperationName {
	return repository_query.IntMoreThan
}

func (BoltIntMoreThan) Apply(a reflect.Value, b reflect.Value) bool {
	return a.Int() > b.Int()
}

type BoltIntLessThan struct{}

func (BoltIntLessThan) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) {
		return nil
	}
	return errors.New("needs an integer")
}

func (BoltIntLessThan) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) {
		return nil
	}
	return errors.New("needs an integer")
}

func (BoltIntLessThan) GetName() repository_query.FilterOperationName {
	return repository_query.IntLessThan
}

func (BoltIntLessThan) Apply(a reflect.Value, b reflect.Value) bool {
	return a.Int() < b.Int()
}

type BoltIntEqual struct{}

func (BoltIntEqual) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) {
		return nil
	}
	return errors.New("needs an integer")
}

func (BoltIntEqual) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) {
		return nil
	}
	return errors.New("needs an integer")
}

func (BoltIntEqual) GetName() repository_query.FilterOperationName {
	return repository_query.IntEqual
}

func (BoltIntEqual) Apply(a reflect.Value, b reflect.Value) bool {
	return a.Int() == b.Int()
}

type BoltIntNotEqual struct{}

func (BoltIntNotEqual) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) {
		return nil
	}
	return errors.New("needs an integer")
}

func (BoltIntNotEqual) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeInt(r) {
		return nil
	}
	return errors.New("needs an integer")
}

func (BoltIntNotEqual) GetName() repository_query.FilterOperationName {
	return repository_query.IntNotEqual
}

func (BoltIntNotEqual) Apply(a reflect.Value, b reflect.Value) bool {
	return a.Int() != b.Int()
}

type BoltStringEqual struct{}

func (BoltStringEqual) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeString(r) {
		return nil
	}
	return errors.New("needs an string")
}

func (BoltStringEqual) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeString(r) {
		return nil
	}
	return errors.New("needs an string")
}

func (BoltStringEqual) GetName() repository_query.FilterOperationName {
	return repository_query.StringEqual
}

func (BoltStringEqual) Apply(a reflect.Value, b reflect.Value) bool {
	return strings.ToLower(a.String()) == strings.ToLower(b.String())
}

type BoltStringNotEqual struct{}

func (BoltStringNotEqual) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeString(r) {
		return nil
	}
	return errors.New("needs an string")
}

func (BoltStringNotEqual) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeString(r) {
		return nil
	}
	return errors.New("needs an string")
}

func (BoltStringNotEqual) GetName() repository_query.FilterOperationName {
	return repository_query.StringNotEqual
}

func (BoltStringNotEqual) Apply(a reflect.Value, b reflect.Value) bool {
	return strings.ToLower(a.String()) != strings.ToLower(b.String())
}

type BoltStringHasSubstr struct{}

func (BoltStringHasSubstr) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeString(r) {
		return nil
	}
	return errors.New("needs an string")
}

func (BoltStringHasSubstr) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeString(r) {
		return nil
	}
	return errors.New("needs an string")
}

func (BoltStringHasSubstr) GetName() repository_query.FilterOperationName {
	return repository_query.StringHasSubstr
}

func (BoltStringHasSubstr) Apply(a reflect.Value, b reflect.Value) bool {
	return strings.Contains(strings.ToLower(a.String()), strings.ToLower(b.String()))
}

type BoltFloatMoreThan struct{}

func (BoltFloatMoreThan) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeFloat(r) {
		return nil
	}
	return errors.New("needs an float")
}

func (BoltFloatMoreThan) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeFloat(r) {
		return nil
	}
	return errors.New("needs an float")
}

func (BoltFloatMoreThan) GetName() repository_query.FilterOperationName {
	return repository_query.FloatMoreThan
}

func (BoltFloatMoreThan) Apply(a reflect.Value, b reflect.Value) bool {
	return a.Float() > b.Float()
}

type BoltFloatLessThan struct{}

func (BoltFloatLessThan) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeFloat(r) {
		return nil
	}
	return errors.New("needs an float")
}

func (BoltFloatLessThan) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeFloat(r) {
		return nil
	}
	return errors.New("needs an float")
}

func (BoltFloatLessThan) GetName() repository_query.FilterOperationName {
	return repository_query.FloatLessThan
}

func (BoltFloatLessThan) Apply(a reflect.Value, b reflect.Value) bool {
	return a.Float() < b.Float()
}

type BoltFloatEqual struct{}

func (BoltFloatEqual) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeFloat(r) {
		return nil
	}
	return errors.New("needs an float")
}

func (BoltFloatEqual) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeFloat(r) {
		return nil
	}
	return errors.New("needs an float")
}

func (BoltFloatEqual) GetName() repository_query.FilterOperationName {
	return repository_query.FloatEqual
}

func (BoltFloatEqual) Apply(a reflect.Value, b reflect.Value) bool {
	return a.Float() == b.Float()
}

type BoltFloatNotEqual struct{}

func (BoltFloatNotEqual) CanAcceptModelType(r reflect.Type) error {
	if repository_query.IsTypeFloat(r) {
		return nil
	}
	return errors.New("needs an float")
}

func (BoltFloatNotEqual) CanAcceptUserType(r reflect.Type) error {
	if repository_query.IsTypeFloat(r) {
		return nil
	}
	return errors.New("needs an float")
}

func (BoltFloatNotEqual) GetName() repository_query.FilterOperationName {
	return repository_query.FloatNotEqual
}

func (BoltFloatNotEqual) Apply(a reflect.Value, b reflect.Value) bool {
	return a.Float() != b.Float()
}

type BoltUuidEqual struct {
}

func (BoltUuidEqual) CanAcceptModelType(r reflect.Type) error {
	if reflect.TypeOf(uuid.UUID{}) == r {
		return nil
	}
	return errors.New("needs an uuid")
}

func (BoltUuidEqual) CanAcceptUserType(r reflect.Type) error {
	if reflect.TypeOf(uuid.UUID{}) == r {
		return nil
	}
	return errors.New("needs an uuid")
}

func (BoltUuidEqual) GetName() repository_query.FilterOperationName {
	return "uuid_equal"
}

func (BoltUuidEqual) Apply(a reflect.Value, b reflect.Value) bool {
	return a.Interface().(uuid.UUID) != b.Interface().(uuid.UUID)
}

type BoltIsNull struct{}

func (BoltIsNull) CanAcceptModelType(r reflect.Type) error {
	if r.Kind() == reflect.Map ||
		r.Kind() == reflect.Chan ||
		r.Kind() == reflect.Pointer ||
		r.Kind() == reflect.Func ||
		r.Kind() == reflect.Interface {
		return nil
	}
	return errors.New("needs to be a nillable type")
}

func (BoltIsNull) CanAcceptUserType(r reflect.Type) error {
	return nil
}

func (BoltIsNull) GetName() repository_query.FilterOperationName {
	return repository_query.FloatNotEqual
}

func (BoltIsNull) Apply(a reflect.Value, b reflect.Value) bool {
	return a.IsNil()
}
