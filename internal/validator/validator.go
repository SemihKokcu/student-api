package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// MsgForTag, tag tipine göre özel mesaj döner
func MsgForTag(tag string, param string) string {
	switch tag {
	case "required":
		return "Bu alan zorunludur."
	case "email":
		return "Geçersiz e-posta adresi."
	case "min":
		return fmt.Sprintf("En az %s karakter olmalıdır.", param)
	case "max":
		return fmt.Sprintf("En fazla %s karakter olmalıdır.", param)
	case "gte":
		return fmt.Sprintf("%s değerinden büyük veya eşit olmalıdır.", param)
	case "lte":
		return fmt.Sprintf("%s değerinden küçük veya eşit olmalıdır.", param)
	}
	return "Geçersiz değer."
}

func MapValidationErrors(err error) map[string]string {
	var ve validator.ValidationErrors
	out := make(map[string]string)

	if errors.As(err, &ve) {
		for _, fe := range ve {
			out[fe.Field()] = MsgForTag(fe.Tag(), fe.Param())
		}
	}
	return out
}
