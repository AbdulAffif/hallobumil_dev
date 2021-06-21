package dto
import (
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-ozzo/ozzo-validation/v4"
)
type RegisterDTO struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Height string `json:"height" form:"height"`
	Phone string `json:"phone" form:"phone"`
	Birthdate string `json:"birthdate" form:"birthdate"`
	ExpectedDate string `json:"expected_date" form:"expected_date"`
	ProfilePicture *string `json:"profile_picture" form:"profile_picture"`
	RegisterOsType *string `json:"register_os_type" form:"register_os_type"`
	Identifier string`json:"identifier" form:"identifier"`
}

func (reg RegisterDTO) Validate() error {
	return validation.ValidateStruct(&reg,
		validation.Field(&reg.Name,validation.Required,validation.Length(6, 100)),
		validation.Field(&reg.Email,validation.Required,is.Email),
		validation.Field(&reg.Password,validation.Required,validation.Length(6, 50)),
		validation.Field(&reg.Phone,validation.Required,validation.Length(12, 13)),
		validation.Field(&reg.Birthdate,validation.Required),
	)
}
