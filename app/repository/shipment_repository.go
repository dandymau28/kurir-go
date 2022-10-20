package repository

import (
	"github.com/kurir-go/app/dto"
	"github.com/kurir-go/app/model"
)

type ShipmentRepository interface {
	SaveShipment(shipment *model.Shipment) error
	UpdateShipment(shipment *model.Shipment) error
	GetAllShipment() []dto.ShipmentGetResponse
	GetShipmentByCourier(courier_id uint64) []dto.ShipmentGetResponse
	GetShipmentByID(shipment_id uint64) dto.ShipmentHistoryResponse
	UpdateDeliveryShipment(shipment_id uint64) error
	GetShipmentHistoryByID(shipment_id uint64) []model.ShipmentHistory
	SaveHistory(history *model.ShipmentHistory) error
}
