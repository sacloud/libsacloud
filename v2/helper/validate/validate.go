package validate

import "gopkg.in/go-playground/validator.v9"

// Struct go-playground/validatorを利用してバリデーションを行う
func Struct(s interface{}) error {
	return validator.New().Struct(s)
}
