package model

import (
	"time"

	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	ID              uint64    `json:"id" gorm:"primary_key:auto_increment"`
	UserID          uint64    `json:"user_id"`
	AWBNumber       string    `json:"awb_number" gorm:"column:awb_number;type:varchar(255)"`
	ClientName      string    `json:"client_name" gorm:"column:client_name"`
	ClientAddress   string    `json:"client_address" gorm:"column:client_address"`
	DeliveredTo     string    `json:"receiver" gorm:"column:delivered_to"`
	Status          string    `json:"status" gorm:"column:status;comment:received,on_process,delivered"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeliveredAt     time.Time `json:"delivered_at" gorm:"column:delivered_at"`
	User            User      `gorm:"foreignKey:UserID"`
	ShipmentHistory []ShipmentHistory
}
