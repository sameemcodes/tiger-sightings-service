package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"tigerhall-kittens/cmd/models"
	"tigerhall-kittens/cmd/service"

	"github.com/gin-gonic/gin"
)

type TigerController interface {
	GetTigerById(ctx *gin.Context)
	GetAllTigers(ctx *gin.Context)
	CreateNewTiger(ctx *gin.Context)
	UpdateTiger(ctx *gin.Context)
	DeleteTigerById(ctx *gin.Context)
	CheckIfTigerExists(ctx *gin.Context)
}

type tigerController struct {
	tigerService service.TigerService
}

func NewTigerController(tigerService service.TigerService) TigerController {
	return &tigerController{
		tigerService: tigerService,
	}
}

// GetTigerById godoc
// @Summary Get a tiger by tigerId
// @Tags Tiger-Controller
// @Accept */*
// @Produce json
// @Param tigerId path string true "tigerId"
// @Router /tiger/v1/tiger_id/{tigerId} [get]
func (c *tigerController) GetTigerById(ctx *gin.Context) {
	tigerId := ctx.Params.ByName("tigerId")
	tigerDto, err := c.tigerService.GetTigerById(ctx, tigerId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetTigerById": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, tigerDto)
	}
}

// GetAllTigers godoc
// @Summary Get all tigers
// @Tags Tiger-Controller
// @Accept */*
// @Produce json
// @Param offset query int true "offset"
// @Param limit query int true "limit"
// @Router /tiger/v1/fetch_all [get]
func (c *tigerController) GetAllTigers(ctx *gin.Context) {
	offsetStr := ctx.Query("offset")
	limitStr := ctx.Query("limit")
	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)
	tigerDto, err := c.tigerService.GetAllTigers(ctx, offset, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during GetAllTigers": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, tigerDto)
	}
}

// CreateNewTiger godoc
// @Summary Create a new tiger
// @Tags Tiger-Controller
// @Accept */*
// @Param tiger body models.Tiger true "Tiger body with timestamp format yyyy-mm-dd HH:ii:ss"
// @Success 200
// @Failure 404
// @Failure 500
// @Produce json
// @Router /tiger/v1/create_new [post]
func (c *tigerController) CreateNewTiger(ctx *gin.Context) {
	var tiger models.Tiger
	err := ctx.BindJSON(&tiger)
	if err != nil {
		fmt.Println("Error during CreateNewTiger", err.Error())
	}
	tigerDto, err := c.tigerService.CreateNewTiger(ctx, tiger)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during CreateNewTiger ": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, tigerDto)
	}
}

// UpdateTiger godoc
// @Summary Update a tiger
// @Tags Tiger-Controller
// @Param tiger body models.Tiger true "Tiger"
// @Accept */*
// @Produce json
// @Router /tiger/v1/update [put]
func (c *tigerController) UpdateTiger(ctx *gin.Context) {
	var tiger models.Tiger
	err := ctx.BindJSON(&tiger)
	if err != nil {
		fmt.Println("Error during UpdateTiger", err.Error())
	}
	tigerDto, err := c.tigerService.UpdateTiger(ctx, tiger)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during UpdateTiger": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, tigerDto)
	}
}

// DeleteTigerById godoc
// @Summary Delete a tiger by tigerId
// @Tags Tiger-Controller
// @Accept */*
// @Param tigerId path string true "tigerId"
// @Produce json
// @Success 200
// @Router /tiger/v1/deletebyTigerId/{tigerId} [delete]
func (c *tigerController) DeleteTigerById(ctx *gin.Context) {
	tigerId := ctx.Params.ByName("tigerId")
	err := c.tigerService.DeleteTiger(ctx, tigerId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during DeleteTigerById": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Tiger deleted successfully"})
	}
}

// CheckIfTigerExists godoc
// @Summary Check if a tiger exists by tigerId
// @Tags Tiger-Controller
// @Accept */*
// @Param tigerId path string true "tigerId"
// @Produce json
// @Success 200
// @Router /tiger/v1/checkIfTigerExists/{tigerId} [get]
func (c *tigerController) CheckIfTigerExists(ctx *gin.Context) {
	tigerId := ctx.Params.ByName("tigerId")
	tigerExists, err := c.tigerService.CheckIfTigerExists(ctx, tigerId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error during CheckIfTigerExists": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"tigerExists": tigerExists})
	}
}
