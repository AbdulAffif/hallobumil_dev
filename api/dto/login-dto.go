package dto

import (
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-ozzo/ozzo-validation/v4"
)

type LoginDTOEmail struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Identifier string`json:"identifier" form:"identifier"`
}

func (le LoginDTOEmail) Validate() error {
	return validation.ValidateStruct(&le,
		validation.Field(&le.Email,validation.Required,is.Email),
		validation.Field(&le.Password,validation.Required,validation.Length(6,50)),
		validation.Field(&le.Identifier,validation.Required,validation.Length(1,50)),
		)
}

type LoginDTOPhone struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	Identifier string`json:"identifier" form:"identifier"`
}

func (lp LoginDTOPhone) Validate() error {
	return validation.ValidateStruct(&lp,
		validation.Field(&lp.Phone,validation.Required,validation.Length(12,13)),
		validation.Field(&lp.Password,validation.Required,validation.Length(6,50)),
		validation.Field(&lp.Identifier,validation.Required,validation.Length(1,50)),
		)
}

