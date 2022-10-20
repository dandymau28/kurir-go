package service

import (
	"errors"
	"log"
	"time"

	"github.com/kurir-go/app/dto"
	"github.com/kurir-go/app/model"
	"github.com/kurir-go/app/repository"
)

type ShipmentServiceImpl struct {
	shipmentRepo repository.ShipmentRepository
	userRepo     repository.UserRepository
}

func NewShipmentService(shipmentRepo repository.ShipmentRepository, userRepo repository.UserRepository) ShipmentService {
	return &ShipmentServiceImpl{
		shipmentRepo: shipmentRepo,
		userRepo:     userRepo,
	}
}

func (s *ShipmentServiceImpl) AddShipment(client_name string, client_address string, courier string, awb_no string) error {
	log.Printf("Add Shipment service: start")

	user := s.userRepo.FindByUsername(courier)

	if (user == dto.User{}) {
		log.Printf("Add Shipment service: courier not found")
		return errors.New("no courier found")
	}

	shipment := &model.Shipment{
		AWBNumber:     awb_no,
		ClientName:    client_name,
		ClientAddress: client_address,
		UserID:        user.ID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Status:        "on_process",
	}

	errShipment := s.shipmentRepo.SaveShipment(shipment)

	if errShipment != nil {
		log.Printf("Add Shipment service: failed to save shipment %s", errShipment)
		return errors.New("failed to save shipment")
	}

	history := &model.ShipmentHistory{
		AWBNumber:     shipment.AWBNumber,
		ClientName:    client_name,
		ClientAddress: client_address,
		Username:      user.Username,
		Name:          user.Name,
		Role:          user.Role,
		ShipmentID:    shipment.ID,
		CreatedAt:     shipment.CreatedAt,
		UpdatedAt:     shipment.UpdatedAt,
		Status:        shipment.Status,
	}

	errHistory := s.shipmentRepo.SaveHistory(history)

	if errHistory != nil {
		log.Printf("Add Shipment service: failed to save history %s", errHistory)
	}

	log.Printf("Add Shipment service: finished")
	return nil
}

func (s *ShipmentServiceImpl) UpdateDeliveryShipment(shipment_id uint64) error {
	errUpdate := s.shipmentRepo.UpdateDeliveryShipment(shipment_id)

	if errUpdate != nil {
		log.Printf("UpdateDeliveryShipment: fail update: %s", errUpdate.Error())
		return errUpdate
	}

	return nil
}

func (s *ShipmentServiceImpl) GetAllShipment() []dto.ShipmentGetResponse {
	data := s.shipmentRepo.GetAllShipment()

	return data
}

func (s *ShipmentServiceImpl) GetShipmentByCourier(courier string) ([]dto.ShipmentGetResponse, error) {
	user := s.userRepo.FindByUsername(courier)

	if (user == dto.User{}) {
		log.Printf("GetShipmentByCourier: no courier found")
		return []dto.ShipmentGetResponse{{}}, errors.New("no courier found")
	}

	shipment := s.shipmentRepo.GetShipmentByCourier(user.ID)

	return shipment, nil
}

func (s *ShipmentServiceImpl) GetShipmentDetailByID(shipment_id uint64) dto.ShipmentHistoryResponse {
	shipment := s.shipmentRepo.GetShipmentByID(shipment_id)

	shipment.ShipmentHistory = s.shipmentRepo.GetShipmentHistoryByID(shipment_id)

	return shipment
}
