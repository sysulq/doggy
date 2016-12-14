package doggy

import (
	"gopkg.in/asaskevich/govalidator.v4"
)

// ValidateStruct use tags for fields.
// result will be equal to `false` if there are any errors.
func ValidateStruct(s interface{}) error {
	_, err := govalidator.ValidateStruct(s)
	return err
}
