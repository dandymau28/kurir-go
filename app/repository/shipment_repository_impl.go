package repository

import (
	"github.com/kurir-go/app/dto"
	"github.com/kurir-go/app/model"
	"gorm.io/gorm"
)

type ShipmentRepositoryImpl struct {
	db *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) ShipmentRepository {
	return &ShipmentRepositoryImpl{
		db: db,
	}
}

func (repository *ShipmentRepositoryImpl) SaveShipment(shipment *model.Shipment) error {
	result := repository.db.Create(shipment)
	return result.Error
}
func (repository *ShipmentRepositoryImpl) UpdateShipment(shipment *model.Shipment) error {
	result := repository.db.Save(shipment)
	return result.Error
}
func (repository *ShipmentRepositoryImpl) UpdateDeliveryShipment(shipment_id uint64) error {
	result := repository.db.Model(&model.Shipment{}).Where("id = ?", shipment_id).Update("status", "delivered")
	return result.Error
}
func (repository *ShipmentRepositoryImpl) GetAllShipment() []dto.ShipmentGetResponse {
	var result []dto.ShipmentGetResponse
	repository.db.Model(&model.Shipment{}).Select("name", "client_name", "client_address", "awb_number", "status", "shipments.created_at", "shipments.updated_at", "shipments.id").Joins("left join users on users.id = shipments.user_id").Find(&result)
	return result
}
func (repository *ShipmentRepositoryImpl) GetShipmentByCourier(courier_id uint64) []dto.ShipmentGetResponse {
	var result []dto.ShipmentGetResponse
	repository.db.Model(&model.Shipment{}).Select("shipments.id", "users.name", "client_name", "shipments.created_at", "shipments.updated_at", "awb_number", "status").Joins("left join users on users.id = shipments.user_id").Where("user_id = ?", courier_id).Find(&result)
	return result
}

func (repository *ShipmentRepositoryImpl) SaveHistory(history *model.ShipmentHistory) error {
	result := repository.db.Create(history)
	return result.Error
}

func (repository *ShipmentRepositoryImpl) GetShipmentByID(shipment_id uint64) dto.ShipmentHistoryResponse {
	var result dto.ShipmentHistoryResponse
	repository.db.Model(&model.Shipment{}).Select("id", "name", "role", "client_name", "client_address", "awb_number", "status").Joins("join users on users.id = shipment.user_id").Where("shipment.id = ?", shipment_id).Take(&result)
	return result
}

func (repository *ShipmentRepositoryImpl) GetShipmentHistoryByID(shipment_id uint64) []model.ShipmentHistory {
	var result []model.ShipmentHistory
	repository.db.Where("shipment_id = ?", shipment_id).Find(&result)
	return result
}
