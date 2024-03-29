package service

import (
	"context"
	"fmt"
	"tigerhall-kittens/cmd/models"
	"tigerhall-kittens/cmd/repository"

	"github.com/mitchellh/mapstructure"
)

type TigerService interface {
	GetTigerById(ctx context.Context, tigerId string) (_ *models.Tiger, err error)
	GetAllTigers(ctx context.Context, offset int, limit int) (_ []models.Tiger, err error)
	CreateNewTiger(ctx context.Context, tiger models.Tiger) (_ models.Tiger, err error)
	UpdateTiger(ctx context.Context, tiger models.Tiger) (_ *models.Tiger, err error)
	DeleteTiger(ctx context.Context, tigerId string) (err error)
	CheckIfTigerExists(ctx context.Context, tigerId string) (bool, error)
}

type tigerService struct {
	tigerRepository repository.TigerRepository
}

func NewTigerService(tigerRepo repository.TigerRepository) TigerService {
	return &tigerService{
		tigerRepository: tigerRepo,
	}
}

func (service *tigerService) GetTigerById(ctx context.Context, tigerId string) (_ *models.Tiger, err error) {
	entity, errorDb := service.tigerRepository.GetTigerById(ctx, tigerId)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}

func (service *tigerService) GetAllTigers(ctx context.Context, offset int, limit int) (_ []models.Tiger, err error) {
	entity, errorDb := service.tigerRepository.GetAllTigers(ctx, offset, limit)
	if errorDb != nil {
		return nil, errorDb
	}
	return entity, nil
}

func (service *tigerService) CreateNewTiger(ctx context.Context, tiger models.Tiger) (_ models.Tiger, err error) {
	entity, errorDb := service.tigerRepository.CreateNewTiger(ctx, tiger)
	if errorDb != nil {
		return tiger, errorDb
	}
	return entity, nil
}

func (service *tigerService) UpdateTiger(ctx context.Context, tiger models.Tiger) (_ *models.Tiger, err error) {
	var tigerId = tiger.TigerID
	newTiger, err := service.GetTigerById(ctx, tigerId)
	if err != nil {
		fmt.Println("Get Tiger By Id not found  ", err)
		return &tiger, err
	}

	err2 := mapstructure.Decode(tiger, &newTiger)

	if err2 != nil {
		fmt.Println("err2  ", err2)
		return &tiger, err2
	}
	entity, errorDb := service.tigerRepository.SaveTiger(ctx, newTiger)
	if errorDb != nil {
		return &tiger, errorDb
	}
	return entity, nil
}

func (service *tigerService) DeleteTiger(ctx context.Context, tigerId string) (err error) {
	errorDb := service.tigerRepository.DeleteTigerById(ctx, tigerId)
	if errorDb != nil {
		return errorDb
	}
	return nil
}

func (service *tigerService) CheckIfTigerExists(ctx context.Context, tigerId string) (bool, error) {
	_, errorDb := service.tigerRepository.GetTigerById(ctx, tigerId)
	if errorDb != nil {
		return false, errorDb
	}
	return true, nil
}
