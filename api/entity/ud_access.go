package entity

import "time"

type Ud_access struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Id_user     uint64    `gorm:"type:Integer" json:"id_user"`
	Identifier  string    `gorm:"type:varchar(255)" json:"identifier"`
	Token       string    `gorm:"type:varchar(255)" json:"token"`
	Last_update time.Time `gorm:"type:datetime" json:"last_update"`
	Is_deleted  bool       `json:"is_deleted"`
}
