package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/simpleBank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	currency, ok := fieldLevel.Field().Interface().(string)
	if ok {
		// check currency is supported or not
		return util.IsSupportedCurrency(currency)
	}
	return false
}
