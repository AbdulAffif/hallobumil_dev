package entity

import "time"

type Result struct {
	ID              uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Email           string    `gorm:"type:varchar(255)" json:"email"`
	Is_verified     bool       `json:"is_verified"`
	Status          int       `gorm:"type:integer" json:"status"`
	Name            string    `gorm:"type:varchar(255)" json:"name"`
	ProfilePicture *string    `gorm:"type:varchar(255)" json:"profile_picture"`
	Birthdate       time.Time `gorm:"type:datetime" json:"birthdate"`
	Height          string   `json:"height"`
	Iteration       int       `json:"iteration"`
	Phone           string    `gorm:"type:varchar(25)" json:"phone"`
}

type JsonRegister struct {
	User Result `json:"user"`
	UdChildBirths Ud_pregnancy `json:"ud_child_births"`
	Token string `json:"token"`
	JwtToken string `json:"jwt_token"`
	ServerTime time.Time `json:"server_time"`
}
