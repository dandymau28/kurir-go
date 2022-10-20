package service

import (
	"github.com/kurir-go/app/dto"
)

type ShipmentService interface {
	AddShipment(client_name string, client_address string, courier string, awb_no string) error
	UpdateDeliveryShipment(shipment_id uint64) error
	GetAllShipment() []dto.ShipmentGetResponse
	GetShipmentByCourier(courier string) ([]dto.ShipmentGetResponse, error)
	GetShipmentDetailByID(shipment_id uint64) dto.ShipmentHistoryResponse
}
