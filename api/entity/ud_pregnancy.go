package entity

import (
	"time"
)

type Ud_pregnancy struct {
	ID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	IdUser uint64 `json:"id_user"`
	Iteration int `json:"iteration"`
	Expected_date time.Time `json:"expected_date"`
}
