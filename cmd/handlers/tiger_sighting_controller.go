package controller

import (
	"net/http"
	"strconv"
	"tigerhall-kittens/cmd/models"
	"tigerhall-kittens/cmd/service"

	"github.com/gin-gonic/gin"
)

type TigerSightingController interface {
	GetTigerSightingById(ctx *gin.Context)
	GetAllTigerSightings(ctx *gin.Context)
	CreateNewTigerSighting(ctx *gin.Context)
	UpdateTigerSighting(ctx *gin.Context)
	DeleteTigerSighting(ctx *gin.Context)
	GetTigerSightingsByTigerId(ctx *gin.Context)
}

type tigerSightingController struct {
	tigerSightingService service.TigerSightingService
}

func NewTigerSightingController(tigerSightingService service.TigerSightingService) TigerSightingController {
	return &tigerSightingController{
		tigerSightingService: tigerSightingService,
	}
}

// GetTigerSightingById godoc
// @Summary Get a tiger sighting by sightingId
// @Tags TigerSighting-Controller
// @Accept */*
// @Produce json
// @Param sightingId path string true "sightingId"
// @Router /tigerSighting/v1/sighting_id/{sightingId} [get]
func (c *tigerSightingController) GetTigerSightingById(ctx *gin.Context) {
	sightingId := ctx.Params.ByName("sightingId")
	sightingData, err := c.tigerSightingService.GetTigerSightingById(ctx, sightingId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetTigerSightingById": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, sightingData)
	}
}

// GetAllTigerSightings godoc
// @Summary Get all tiger sightings
// @Tags TigerSighting-Controller
// @Accept */*
// @Produce json
// @Param offset query int true "offset"
// @Param limit query int true "limit"
// @Router /tigerSighting/v1/fetch_all [get]
func (c *tigerSightingController) GetAllTigerSightings(ctx *gin.Context) {
	offsetStr := ctx.Query("offset")
	limitStr := ctx.Query("limit")
	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)
	sightings, err := c.tigerSightingService.GetAllTigerSightings(ctx, offset, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetAllTigerSightings": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, sightings)
	}
}

// CreateNewTigerSighting godoc
// @Summary Create a new tiger sighting
// @Tags TigerSighting-Controller
// @Accept */*
// @Param sightingData body models.TigerSightingData true "TigerSightingData"
// @Success 200
// @Failure 404
// @Failure 500
// @Produce json
// @Router /tigerSighting/v1/create_new [post]
func (c *tigerSightingController) CreateNewTigerSighting(ctx *gin.Context) {
	var sightingData models.TigerSightingData
	if err := ctx.BindJSON(&sightingData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sightingData, err := c.tigerSightingService.CreateNewTigerSighting(ctx, sightingData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during CreateNewTigerSighting ": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, sightingData)
	}
}

// UpdateTigerSighting godoc
// @Summary Update a tiger sighting
// @Tags TigerSighting-Controller
// @Param sightingData body models.TigerSightingData true "TigerSightingData"
// @Accept */*
// @Produce json
// @Router /tigerSighting/v1/update [put]
func (c *tigerSightingController) UpdateTigerSighting(ctx *gin.Context) {
	var sightingData models.TigerSightingData
	if err := ctx.BindJSON(&sightingData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSightingData, err := c.tigerSightingService.UpdateTigerSighting(ctx, sightingData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during UpdateTigerSighting": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedSightingData)
}

// DeleteTigerSighting godoc
// @Summary Delete a tiger sighting by sightingId
// @Tags TigerSighting-Controller
// @Param sightingId path string true "sightingId"
// @Accept */*
// @Produce json
// @Success 200
// @Router /tigerSighting/v1/deletebySightingId/{sightingId} [delete]
func (c *tigerSightingController) DeleteTigerSighting(ctx *gin.Context) {
	sightingId := ctx.Params.ByName("sightingId")
	err := c.tigerSightingService.DeleteTigerSighting(ctx, sightingId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during DeleteTigerSighting": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Tiger sighting deleted successfully"})
	}
}

// GetTigerSightingsByTigerId godoc
// @Summary Get all tiger sightings for a tiger
// @Tags TigerSighting-Controller
// @Param tigerId path string true "tigerId"
// @Param offset query int true "offset"
// @Param limit query int true "limit"
// @Accept */*
// @Produce json
// @Router /tigerSighting/v1/tiger_id/{tigerId} [get]
func (c *tigerSightingController) GetTigerSightingsByTigerId(ctx *gin.Context) {
	offsetStr := ctx.Query("offset")
	limitStr := ctx.Query("limit")
	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)
	tigerId := ctx.Params.ByName("tigerId")
	sightings, err := c.tigerSightingService.GetTigerSightingsByTigerId(ctx, tigerId, offset, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetTigerSightingsByTigerId": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, sightings)
	}
}