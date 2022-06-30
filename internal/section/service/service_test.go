package service

import (
	"context"
	"errors"
	"testing"

	"github.com/cyruzin/meli-frescos/internal/section/domain"
	"github.com/cyruzin/meli-frescos/internal/section/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStore(t *testing.T) {
	mockSectionsRepo := new(mocks.SectionRepository)
	mockSection := &domain.Section{
		ID:                 1,
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
		mockSectionsRepo.On("Store",
			mock.Anything,
			mock.Anything,
		).Return(mockSection, nil).Once()

		s := NewSection(mockSectionsRepo)

		sec, err := s.Store(context.TODO(), mockSection)

		expectedCurrentCapacity := int64(500)

		assert.NoError(t, err)
		assert.Equal(t, expectedCurrentCapacity, sec.CurrentCapacity)

		mockSectionsRepo.AssertExpectations(t)
	})

	t.Run("fail", func(t *testing.T) {
		mockSectionsRepo.On("Store",
			mock.Anything,
			mock.Anything,
		).Return(&domain.Section{}, errors.New("failed to store")).Once()

		s := NewSection(mockSectionsRepo)

		_, err := s.Store(context.TODO(), mockSection)

		assert.Error(t, err)

		mockSectionsRepo.AssertExpectations(t)
	})
}
