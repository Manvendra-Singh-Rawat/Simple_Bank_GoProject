package utils

import "github.com/go-playground/validator/v10"

var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	currency, ok := fieldLevel.Field().Interface().(string)
	if ok {
		// Check currency is supported or not
		return isSupportedCurrency(currency)
	}
	return false
}
