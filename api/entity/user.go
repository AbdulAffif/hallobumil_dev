package entity

import (
	"time"
)

type User struct {
	ID                                 uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Email                              string    `gorm:"type:varchar(255),unique;not null" json:"email"`
	IdFacebook                        *string    `gorm:"type:varchar(255)" json:"id_facebook"`
	Password                           string    `gorm:"->;<-:create" json:"-"`
	New_password                       *string    `gorm:"->;<-:create" json:"-"`
	Status                             int       `gorm:"type:integer" json:"status"`
	Created_date                       time.Time `gorm:"type:datetime" json:"created_date"`
	Last_update                        time.Time `gorm:"type:datetime" json:"last_update"`
	Is_verified                        bool       `json:"is_verified"`
	Verification_code                  string    `gorm:"type:varchar(255)" json:"verification_code"`
	Is_deleted                         bool       `json:"is_deleted"`
	Is_new                             bool       `json:"is_new"`
	RegisterOsType                    *string    `json:"register_os_type"`
	AppleId                           *string    `gorm:"type:varchar(255)" json:"apple_id"`
	Temp_email                         *string    `gorm:"type:varchar(255)" json:"temp_email"`
	Otp_verifivation_code_created_date time.Time `gorm:"type:datetime" json:"otp_verifivation_code_created_date"`
	Phone                              string    `gorm:"type:varchar(16)" json:"phone"`
	Temp_phone                         *string    `gorm:"type:varchar(16)" json:"temp_phone"`
	Phone_otp_code                     *string    `gorm:"type:varchar(6)" json:"phone_otp_code"`
	Phone_otp_code_created_at          time.Time `gorm:"type:datetime" json:"phone_otp_code_created_at"`
	Has_verified_phone_number          *bool       `json:"has_verified_phone_number"`
}
