package controller

import (
	"net/http"
	"strconv"

	"github.com/cyruzin/meli-frescos/internal/section/domain"
	"github.com/gin-gonic/gin"
)

type SectionControler struct {
	service domain.SectionService
}

type AppError struct {
	Message string
	Code    int
}

type requestCreate struct {
	SectionNumber      int64 `json:"section_number" binding:"required"`
	CurrentTemperature int16 `json:"current_temperature" binding:"required"`
	MinimumTemperature int16 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int64 `json:"current_capacity" binding:"required"`
	MinimumCapacity    int64 `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int64 `json:"maximum_capacity" binding:"required"`
	WarehouseID        int64 `json:"warehouse_id" binding:"required"`
	ProductTypeID      int64 `json:"product_type_id" binding:"required"`
}

type requestUpdate struct {
	SectionNumber      int64 `json:"section_number"`
	CurrentTemperature int16 `json:"current_temperature"`
	MinimumTemperature int16 `json:"minimum_temperature"`
	CurrentCapacity    int64 `json:"current_capacity"`
	MinimumCapacity    int64 `json:"minimum_capacity"`
	MaximumCapacity    int64 `json:"maximum_capacity"`
	WarehouseID        int64 `json:"warehouse_id"`
	ProductTypeID      int64 `json:"product_type_id"`
}

func NewSectionControler(ctx *gin.Engine, service domain.SectionService) {
	sc := &SectionControler{service: service}

	sr := ctx.Group("/api/v1/sections")
	{
		sr.GET("/", sc.GetAll())
		sr.GET("/:id", sc.GetById())
		sr.POST("/", sc.Post())
		sr.PATCH("/:id", sc.Patch())
		sr.DELETE("/:id", sc.Delete())
	}
}

func (c SectionControler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sections, err := c.service.GetAll(ctx.Request.Context())

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, sections)
	}
}

func (c SectionControler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		section, err := c.service.GetByID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, section)
	}
}

func (c SectionControler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestCreate
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		section, err := c.service.Store(ctx, &domain.Section{
			SectionNumber:      req.SectionNumber,
			CurrentTemperature: req.CurrentTemperature,
			MinimumTemperature: req.MinimumTemperature,
			CurrentCapacity:    req.CurrentCapacity,
			MinimumCapacity:    req.MinimumCapacity,
			MaximumCapacity:    req.MaximumCapacity,
			WarehouseID:        req.WarehouseID,
			ProductTypeID:      req.ProductTypeID,
		})

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, section)
	}
}

func (c SectionControler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		var req requestUpdate
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		section, err := c.service.Update(ctx, &domain.Section{
			ID:                 id,
			SectionNumber:      req.SectionNumber,
			CurrentTemperature: req.CurrentTemperature,
			MinimumTemperature: req.MinimumTemperature,
			CurrentCapacity:    req.CurrentCapacity,
			MinimumCapacity:    req.MinimumCapacity,
			MaximumCapacity:    req.MaximumCapacity,
			WarehouseID:        req.WarehouseID,
			ProductTypeID:      req.ProductTypeID,
		})
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, section)
	}
}

func (c SectionControler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}
