package types

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"

	"gitea.alchemymagic.app/snap/go-common/erroy"
)

type (
	Validator struct {
		*validator.Validate
	}
	tValidatorCustomType struct {
		Func      validator.CustomTypeFunc
		SampleObj any
	}
)

var (
	DefaultValidator      = NewSingleton(func() Validator { return NewValidator() })
	vValidatorCustomTypes []tValidatorCustomType
)

func NewValidator() Validator {
	v := Validator{validator.New()}
	v.init()
	return v
}

func (v Validator) init() {
	v.RegisterCustomTypeFunc(v.extractDecimal, decimal.Decimal{})
	v.RegisterCustomTypeFunc(v.extractNullInt64, NullInt64{})
	v.RegisterCustomTypeFunc(v.extractTimeDuration, TimeDuration{})
	for _, t := range vValidatorCustomTypes {
		v.RegisterCustomTypeFunc(t.Func, t.SampleObj)
	}
}

func (v Validator) extractDecimal(field reflect.Value) any {
	value := field.Interface().(decimal.Decimal)
	if value.Equal(value.Truncate(0)) {
		return value.IntPart()
	}
	valueFloat, _ := value.Float64()
	return valueFloat
}

func (v Validator) extractNullInt64(field reflect.Value) any {
	value := field.Interface().(NullInt64)
	if !value.Valid {
		return nil
	}
	return value.Int64
}

func (v Validator) extractTimeDuration(field reflect.Value) any {
	value := field.Interface().(TimeDuration)
	return value.Duration
}

func ValidatorRegisterCustomType(fn validator.CustomTypeFunc, sampleObj any) {
	vValidatorCustomTypes = append(vValidatorCustomTypes, tValidatorCustomType{
		Func:      fn,
		SampleObj: sampleObj,
	})
}

func ValidateStruct(value any) error {
	if err := DefaultValidator.Get().Struct(value); err != nil {
		return erroy.WrapStack(err, "validate struct")
	}
	return nil
}

func ValidateValue(value any, tag string) error {
	if err := DefaultValidator.Get().Var(value, tag); err != nil {
		return erroy.WrapStack(err, "validate "+tag+" value ")
	}
	return nil
}

func ValidateEmail(email string) error {
	return ValidateValue(email, "email")
}
