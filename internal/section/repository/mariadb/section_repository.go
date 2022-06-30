package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cyruzin/meli-frescos/internal/section/domain"
)

type mariadbRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) domain.SectionRepository {
	return mariadbRepository{db: db}
}

func (m mariadbRepository) GetAll(ctx context.Context) (*[]domain.Section, error) {
	sections := []domain.Section{}

	rows, err := m.db.QueryContext(ctx, "SELECT * FROM sections")
	if err != nil {
		return &sections, err
	}

	defer rows.Close()

	for rows.Next() {
		var section domain.Section

		if err := rows.Scan(
			&section.ID,
			&section.SectionNumber,
			&section.CurrentTemperature,
			&section.MinimumTemperature,
			&section.CurrentCapacity,
			&section.MinimumCapacity,
			&section.MaximumCapacity,
			&section.WarehouseID,
			&section.ProductTypeID,
		); err != nil {
			return &sections, err
		}

		sections = append(sections, section)
	}

	return &sections, nil
}

func (m mariadbRepository) GetByID(ctx context.Context, id int64) (*domain.Section, error) {
	row := m.db.QueryRowContext(ctx, "SELECT * FROM sections WHERE ID = ?", id)

	section := domain.Section{}

	err := row.Scan(
		&section.ID,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseID,
		&section.ProductTypeID,
	)
	// ID not found
	if errors.Is(err, sql.ErrNoRows) {
		return &section, domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return &section, err
	}

	return &section, nil
}

func (m mariadbRepository) Store(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	newSection := domain.Section{}

	query := `INSERT INTO sections 
	(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, 
		maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := m.db.ExecContext(
		ctx,
		query,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseID,
		&section.ProductTypeID,
	)
	if err != nil {
		return &newSection, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return &newSection, err
	}

	section.ID = lastID

	return section, nil
}

func (m mariadbRepository) Update(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	newSection := domain.Section{}

	query := `UPDATE sections SET 
	section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, 
	minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type_id=? WHERE id=?`

	result, err := m.db.ExecContext(
		ctx,
		query,
		&section.SectionNumber,
		&section.CurrentTemperature,
		&section.MinimumTemperature,
		&section.CurrentCapacity,
		&section.MinimumCapacity,
		&section.MaximumCapacity,
		&section.WarehouseID,
		&section.ProductTypeID,
		&section.ID,
	)
	if err != nil {
		return &newSection, err
	}

	affectedRows, err := result.RowsAffected()
	// ID not found
	if affectedRows == 0 {
		return &newSection, domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return &newSection, err
	}

	return section, nil
}

func (m mariadbRepository) Delete(ctx context.Context, id int64) error {
	result, err := m.db.ExecContext(ctx, "DELETE FROM sections WHERE id=?", id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	// ID not found
	if affectedRows == 0 {
		return domain.ErrIDNotFound
	}

	// Other errors
	if err != nil {
		return err
	}

	return nil
}
