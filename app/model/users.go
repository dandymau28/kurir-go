package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64    `json:"id" gorm:"primary_key:auto_increment"`
	Username  string    `json:"username" gorm:"type:varchar(255);column:username"`
	Name      string    `json:"name" gorm:"type:varchar(255);column:name"`
	Role      string    `json:"role" gorm:"comment:kurir,admin;column:role"`
	Password  string    `json:"-" gorm:"->;<-;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	Shipments []Shipment
}
