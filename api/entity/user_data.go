package entity

import (
	"time"
)

type User_data struct {
	ID              uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Id_user         uint64    `json:"id_user"`
	Name            string    `gorm:"type:varchar(255)" json:"name"`
	Birthdate       time.Time `gorm:"type:datetime" json:"birthdate"`
	Height          string   `json:"height"`
	Phone           string    `gorm:"type:varchar(25)" json:"phone"`
	ProfilePicture *string    `gorm:"type:varchar(255)" json:"profile_picture"`
	Last_update     time.Time `gorm:"type:datetime" json:"last_update"`
	Is_deleted      bool       `json:"is_deleted"`
	Iteration       int       `json:"iteration"`
	Is_migrate      bool       `json:"is_migrate"`
	Current_state   int       `json:"current_state"`
	Iteration_pre   int       `json:"iteration_pre"`
	IdCity         *int       `json:"id_city,string,omitempty"`
	Iteration_pasca int       `json:"iteration_pasca"`
	Address         string    `gorm:"type:text" json:"address"`
	IdProvinsi     *int       `json:"id_provinsi,string,omitempty"`
	IdKecamatan    *int       `json:"id_kecamatan,string,omitempty"`
	IdKelurahan    *int       `json:"id_kelurahan,string,omitempty"`
	KodePos        *string    `gorm:"type:varchar(10)" json:"kode_pos"`
}
