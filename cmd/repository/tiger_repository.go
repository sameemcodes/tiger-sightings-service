package repository

import (
	"context"
	"tigerhall-kittens/cmd/models"

	"tigerhall-kittens/cmd/constants"

	"gorm.io/gorm"
)

type TigerRepository interface {
	GetTigerById(ctx context.Context, tigerId string) (_ *models.Tiger, err error)
	GetAllTigers(ctx context.Context, offset int, limit int) (_ []models.Tiger, err error)
	CreateNewTiger(ctx context.Context, tiger models.Tiger) (_ models.Tiger, err error)
	SaveTiger(ctx context.Context, tiger *models.Tiger) (_ *models.Tiger, err error)
	DeleteTigerById(ctx context.Context, tigerId string) (err error)
	CheckIfTigerExists(ctx context.Context, tigerId string) (bool, error)
}

type tigerRepository struct {
	db *gorm.DB // database connection
}

// NewTigerRepository creates a new instance of TigerRepository
func NewTigerRepository(db *gorm.DB) TigerRepository {
	return &tigerRepository{
		db: db,
	}
}

func (tRepo *tigerRepository) GetTigerById(ctx context.Context, tigerId string) (_ *models.Tiger, err error) {
	var tiger models.Tiger
	var dbWithCtx = tRepo.db.WithContext(ctx)
	getTiger := dbWithCtx.Where(constants.WhereTigerId, tigerId).Take(&tiger)
	return &tiger, getTiger.Error
}

func (tRepo *tigerRepository) CreateNewTiger(ctx context.Context, tiger models.Tiger) (_ models.Tiger, err error) {
	var dbWithCtx = tRepo.db.WithContext(ctx)
	createTiger := dbWithCtx.Create(&tiger)
	return tiger, createTiger.Error
}

func (tRepo *tigerRepository) GetAllTigers(ctx context.Context, offset int, limit int) (_ []models.Tiger, err error) {
	var tigers []models.Tiger
	var dbWithCtx = tRepo.db.WithContext(ctx)
	getTigers := dbWithCtx.Order("last_seen_timestamp desc").Offset(offset).Limit(limit).Find(&tigers)
	return tigers, getTigers.Error
}

func (tRepo *tigerRepository) SaveTiger(ctx context.Context, tiger *models.Tiger) (_ *models.Tiger, err error) {
	var dbWithCtx = tRepo.db.WithContext(ctx)
	saveTiger := dbWithCtx.Where(constants.WhereTigerId, tiger.TigerID).Updates(tiger)
	return tiger, saveTiger.Error
}

func (tRepo *tigerRepository) DeleteTigerById(ctx context.Context, tigerId string) (err error) {
	var tiger models.Tiger
	var dbWithCtx = tRepo.db.WithContext(ctx)
	deleteTiger := dbWithCtx.Where(constants.WhereTigerId, tigerId).Delete(&tiger)
	return deleteTiger.Error
}

// check if a tiger exists in the database and name the method CheckIfTigerExists and return boolean
func (tRepo *tigerRepository) CheckIfTigerExists(ctx context.Context, tigerId string) (bool, error) {
	var tiger models.Tiger
	var dbWithCtx = tRepo.db.WithContext(ctx)
	getTiger := dbWithCtx.Where(constants.WhereTigerId, tigerId).Take(&tiger)
	if getTiger.Error != nil {
		return false, getTiger.Error
	}
	return true, nil
}
