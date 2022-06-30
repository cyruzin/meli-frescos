package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cyruzin/meli-frescos/internal/section/domain"
	"github.com/cyruzin/meli-frescos/internal/section/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStore(t *testing.T) {
	mockSection := &domain.Section{
		SectionNumber:      12,
		CurrentTemperature: 10,
		MinimumTemperature: 2,
		CurrentCapacity:    500,
		MinimumCapacity:    10,
		MaximumCapacity:    890,
		WarehouseID:        23,
		ProductTypeID:      56,
	}

	t.Run("success", func(t *testing.T) {
		sectionsServiceMock := new(mocks.SectionService)

		sectionsServiceMock.On("Store",
			mock.Anything,
			mock.Anything,
		).Return(mockSection, nil)

		payload, err := json.Marshal(mockSection)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sections", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		NewSectionControler(engine, sectionsServiceMock)

		sc := SectionControler{service: sectionsServiceMock}

		engine.POST("/api/v1/sections", sc.Post())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})

	t.Run("fail with bad request", func(t *testing.T) {
		sectionsServiceMock := new(mocks.SectionService)
		mockSectionBad := &domain.Section{}

		sectionsServiceMock.On("Store",
			mock.Anything,
			mock.Anything,
		).Return(mockSection, errors.New("bad request"))

		payload, err := json.Marshal(mockSectionBad)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sections", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		NewSectionControler(engine, sectionsServiceMock)

		sc := SectionControler{service: sectionsServiceMock}

		engine.POST("/api/v1/sections", sc.Post())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("fail with internal error", func(t *testing.T) {
		sectionsServiceMock := new(mocks.SectionService)
		sectionsServiceMock.On("Store",
			mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("internal error"))

		payload, err := json.Marshal(mockSection)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/sections", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()

		_, engine := gin.CreateTestContext(rec)

		NewSectionControler(engine, sectionsServiceMock)

		sc := SectionControler{service: sectionsServiceMock}

		engine.POST("/api/v1/sections", sc.Post())

		engine.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		sectionsServiceMock.AssertExpectations(t)
	})
}
