package service

import (
	"context"
	"fmt"
	"tigerhall-kittens/cmd/models"
	"tigerhall-kittens/cmd/repository"

	durable "tigerhall-kittens/cmd/durables"
	"tigerhall-kittens/cmd/utils"

	"github.com/mitchellh/mapstructure"
)

type TigerSightingService interface {
	GetTigerSightingById(ctx context.Context, sightingId string) (_ *models.TigerSightingData, err error)
	GetAllTigerSightings(ctx context.Context, offset int, limit int) (_ []models.TigerSightingData, err error)
	CreateNewTigerSighting(ctx context.Context, sightingData models.TigerSightingData) (_ models.TigerSightingData, err error)
	UpdateTigerSighting(ctx context.Context, sightingData models.TigerSightingData) (_ *models.TigerSightingData, err error)
	DeleteTigerSighting(ctx context.Context, sightingId string) (err error)
	GetTigerSightingsByTigerId(ctx context.Context, tigerId string, offset int, limit int) (_ []models.TigerSightingData, err error)
}

type tigerSightingService struct {
	tigerSightingRepository repository.TigerSightingRepository
}

// NewTigerSightingService creates a new instance of TigerSightingService
func NewTigerSightingService(tsRepo repository.TigerSightingRepository) TigerSightingService {
	return &tigerSightingService{
		tigerSightingRepository: tsRepo,
	}
}

func (service *tigerSightingService) GetTigerSightingById(ctx context.Context, sightingId string) (_ *models.TigerSightingData, err error) {
	entity, errorDb := service.tigerSightingRepository.GetTigerSightingById(ctx, sightingId)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}

func (service *tigerSightingService) GetAllTigerSightings(ctx context.Context, offset int, limit int) (_ []models.TigerSightingData, err error) {
	entity, errorDb := service.tigerSightingRepository.GetAllTigerSightings(ctx, offset, limit)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}

/*
Description: CreateNewTigerSighting
- if the tigerId is not found in the tiger table, create a new entry over the tigertable
- create a new tiger sighting and handle the error if tiger spotted within 5 kms of the last sighting
*/
func (service *tigerSightingService) CreateNewTigerSighting(ctx context.Context, sightingData models.TigerSightingData) (_ models.TigerSightingData, err error) {
	// Check if the tigerId is present in the tiger table
	tigerId := sightingData.TigerID
	tigerService := NewTigerService(repository.NewTigerRepository(durable.MysqlDb))
	tiger, err := tigerService.CheckIfTigerExists(ctx, tigerId)
	if err != nil {
		fmt.Print("err  ", err)
		return models.TigerSightingData{}, err
	}
	if !tiger {
		fmt.Print("Unknown Tiger Sighted .Please create an Entry in Tiger Database  ", tiger)
		return models.TigerSightingData{}, err
	}

	// check if the tiger is spotted within 5 kms of the last sighting
	// take the last sighting of the tiger
	tigerSighting, err := service.GetTigerSightingsByTigerId(ctx, tigerId, 0, 1)
	if err != nil {
		fmt.Print("Tigers not sighted before  ", err)
		return models.TigerSightingData{}, err
	}
	if len(tigerSighting) > 0 {
		tigerSightingLast := tigerSighting[0]
		// check if the tiger is spotted within 5 kms of the last sighting
		tigerSightingLastLat := tigerSightingLast.Latitude
		tigerSightingLastLon := tigerSightingLast.Longitude

		distanceCovered := utils.Haversine(tigerSightingLastLat, tigerSightingLastLon, sightingData.Latitude, sightingData.Longitude)
		fmt.Print("distanceCovered  ", distanceCovered)
		if distanceCovered < 5 {
			fmt.Print("Tiger spotted within 5 kms of the last sighting  ", distanceCovered)
			return models.TigerSightingData{}, err
		}
	}

	entity, errorDb := service.tigerSightingRepository.CreateNewTigerSighting(ctx, sightingData)
	// update the last seen data to the tiger table

	//update only the latest signthing of lat and long and timestamp of input 	UpdateTiger(ctx context.Context, tiger models.Tiger) (_ *models.Tiger, err error)

	updatedTigerModel := models.Tiger{
		TigerID:                tigerId,
		LastSeenCoordinatesLat: sightingData.Latitude,
		LastSeenCoordinatesLon: sightingData.Longitude,
		LastSeenTimestamp:      sightingData.Timestamp,
	}
	fmt.Print("updatedTigerModel  ", updatedTigerModel)

	tigerService.UpdateTiger(ctx, updatedTigerModel)

	if errorDb != nil {
		return sightingData, errorDb
	}
	return entity, nil
}

func (service *tigerSightingService) UpdateTigerSighting(ctx context.Context, sightingData models.TigerSightingData) (_ *models.TigerSightingData, err error) {
	var sightingId = sightingData.SightingID
	newSighting, err := service.GetTigerSightingById(ctx, sightingId)
	fmt.Print("newSighting  ", newSighting, "err ", err)
	// Bind the JSON request body to the sightingData object
	fmt.Print("sightingData ", sightingData)
	err2 := mapstructure.Decode(sightingData, &newSighting)
	fmt.Print("newSighting", newSighting, "sightingData ", sightingData)
	if err2 != nil {
		fmt.Print("err2  ", err2)
		return &sightingData, err2
	}
	entity, errorDb := service.tigerSightingRepository.SaveTigerSighting(ctx, newSighting)
	fmt.Print("entity  ", entity, "errorDb ", errorDb)
	if errorDb != nil {
		return &sightingData, errorDb
	}
	return entity, nil
}

func (service *tigerSightingService) DeleteTigerSighting(ctx context.Context, sightingId string) (err error) {
	// Delete TigerSighting by ID
	errorDb := service.tigerSightingRepository.DeleteTigerSightingById(ctx, sightingId)
	if errorDb != nil {
		return errorDb
	}
	return nil
}

func (service *tigerSightingService) GetTigerSightingsByTigerId(ctx context.Context, tigerId string, offset int, limit int) (_ []models.TigerSightingData, err error) {
	entity, errorDb := service.tigerSightingRepository.GetTigerSightingsByTigerId(ctx, tigerId, offset, limit)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}