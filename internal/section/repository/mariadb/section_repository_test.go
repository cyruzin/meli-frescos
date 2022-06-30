package mariadb

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cyruzin/meli-frescos/internal/section/domain"
	"github.com/stretchr/testify/assert"
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

	query := `INSERT INTO sections 
(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity,
	maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentTemperature,
				mockSection.MinimumTemperature,
				mockSection.CurrentCapacity,
				mockSection.MinimumCapacity,
				mockSection.MaximumCapacity,
				mockSection.WarehouseID,
				mockSection.ProductTypeID,
			).WillReturnResult(sqlmock.NewResult(1, 1)) // last id, // rows affected

		sectionsRepo := NewMariaDBRepository(db)

		sec, err := sectionsRepo.Store(context.TODO(), mockSection)
		assert.NoError(t, err)

		expectedCurrentCapacity := int64(500)

		assert.Equal(t, expectedCurrentCapacity, sec.CurrentCapacity)
	})

	t.Run("failed to store", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(0, 0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1)) // last id, // rows affected

		sectionsRepo := NewMariaDBRepository(db)
		_, err = sectionsRepo.Store(context.TODO(), mockSection)

		assert.Error(t, err)
	})
}
