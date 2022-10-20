package dto

import (
	"time"

	"github.com/kurir-go/app/model"
)

type AddShipmentRequest struct {
	ClientName    string `json:"client_name"`
	ClientAddress string `json:"client_address"`
	AWBNo         string `json:"awb_no"`
}

type ShipmentGetResponse struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	ClientName string    `json:"client_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	AWBNumber  string    `json:"awb_number"`
	Status     string    `json:"status"`
}

type ShipmentHistoryResponse struct {
	ID              uint64                  `json:"id"`
	Name            string                  `json:"courier_name"`
	Role            string                  `json:"role"`
	ClientName      string                  `json:"client_name"`
	ClientAddress   string                  `json:"client_address"`
	AWBNumber       string                  `json:"awb_number"`
	Status          string                  `json:"status"`
	ShipmentHistory []model.ShipmentHistory `json:"shipment_history"`
}
