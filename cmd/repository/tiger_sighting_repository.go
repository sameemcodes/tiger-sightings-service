package repository

import (
	"context"
	"fmt"
	"tigerhall-kittens/cmd/models"

	"tigerhall-kittens/cmd/constants"

	"gorm.io/gorm"
)

type TigerSightingRepository interface {
	GetTigerSightingById(ctx context.Context, sightingId string) (_ *models.TigerSightingData, err error)
	GetAllTigerSightings(ctx context.Context, offset int, limit int) (_ []models.TigerSightingData, err error)
	CreateNewTigerSighting(ctx context.Context, sightingData models.TigerSightingData) (_ models.TigerSightingData, err error)
	SaveTigerSighting(ctx context.Context, sightingData *models.TigerSightingData) (_ *models.TigerSightingData, err error)
	DeleteTigerSightingById(ctx context.Context, sightingId string) (err error)
	GetTigerSightingsByTigerId(ctx context.Context, tigerId string, offset int, limit int) (_ []models.TigerSightingData, err error)
	GetUserSightingsListByTigerId(ctx context.Context, tigerId string) (_ []string, err error)
}

type tigerSightingRepository struct {
	db *gorm.DB // database connection
}

// NewTigerSightingRepository creates a new instance of TigerSightingRepository
func NewTigerSightingRepository(db *gorm.DB) TigerSightingRepository {
	return &tigerSightingRepository{
		db: db,
	}

}

// GetTigerSightingById implements TigerSightingRepository.
func (tsRepo *tigerSightingRepository) GetTigerSightingById(ctx context.Context, sightingId string) (_ *models.TigerSightingData, err error) {
	var tigerSitingModel models.TigerSightingData
	var dbWithCtx = tsRepo.db.WithContext(ctx)
	tigerSightingsByTigerId := dbWithCtx.Where(constants.WhereSightingId, sightingId).Take(&tigerSitingModel)
	return &tigerSitingModel, tigerSightingsByTigerId.Error
}

// GetUserSightingsListByTigerId implements TigerSightingRepository.
func (tsRepo *tigerSightingRepository) GetUserSightingsListByTigerId(ctx context.Context, tigerId string) (_ []string, err error) {
	//GetDistinctUserIDsByTigerID which returns list of string
	var dbWithCtx = tsRepo.db.WithContext(ctx)
	var userIDs []string
	rawQuery := fmt.Sprintf("SELECT DISTINCT user_id FROM tiger_sighting_data WHERE tiger_id = '%s'", tigerId)
	queryResult := dbWithCtx.Raw(rawQuery).Scan(&userIDs)
	return userIDs, queryResult.Error
}

// GetDistinctUserIDsByTigerID retrieves a list of distinct user IDs associated with sightings for a specific tiger.
func (tsRepo *tigerSightingRepository) GetDistinctUserIDsByTigerID(ctx context.Context, tigerID string) (_ []string, err error) {
	var userIDs []string
	var dbWithCtx = tsRepo.db.WithContext(ctx)

	// Replace 'your_table_name' with the actual name of your table
	rawQuery := fmt.Sprintf("SELECT DISTINCT user_id FROM your_table_name WHERE tiger_id = '%s'", tigerID)

	// Execute the raw SQL query
	queryResult := dbWithCtx.Raw(rawQuery).Scan(&userIDs)
	return userIDs, queryResult.Error
}

func (tsRepo *tigerSightingRepository) CreateNewTigerSighting(ctx context.Context, sightingData models.TigerSightingData) (_ models.TigerSightingData, err error) {
	var dbWithCtx = tsRepo.db.WithContext(ctx)
	createSighting := dbWithCtx.Create(&sightingData)
	return sightingData, createSighting.Error
}

func (tsRepo *tigerSightingRepository) GetAllTigerSightings(ctx context.Context, offset int, limit int) (_ []models.TigerSightingData, err error) {
	var sightings []models.TigerSightingData
	var dbWithCtx = tsRepo.db.WithContext(ctx)
	getSightings := dbWithCtx.Offset(offset).Limit(limit).Find(&sightings)
	return sightings, getSightings.Error
}

func (tsRepo *tigerSightingRepository) SaveTigerSighting(ctx context.Context, sightingData *models.TigerSightingData) (_ *models.TigerSightingData, err error) {
	var dbWithCtx = tsRepo.db.WithContext(ctx)
	saveSighting := dbWithCtx.Where(constants.WhereSightingId, sightingData.SightingID).Updates(sightingData)
	return sightingData, saveSighting.Error
}

func (tsRepo *tigerSightingRepository) DeleteTigerSightingById(ctx context.Context, sightingId string) (err error) {
	var sightingData models.TigerSightingData
	var dbWithCtx = tsRepo.db.WithContext(ctx)
	deleteSighting := dbWithCtx.Where(constants.WhereSightingId, sightingId).Delete(&sightingData)
	return deleteSighting.Error
}

func (tsRepo *tigerSightingRepository) GetTigerSightingsByTigerId(ctx context.Context, tigerId string, offset int, limit int) (_ []models.TigerSightingData, err error) {
	var sightings []models.TigerSightingData
	var dbWithCtx = tsRepo.db.WithContext(ctx)
	getSightings := dbWithCtx.Where(constants.WhereTigerId, tigerId).Order("timestamp desc").Offset(offset).Limit(limit).Find(&sightings)
	fmt.Println("getSightings ", getSightings)
	return sightings, getSightings.Error
}
