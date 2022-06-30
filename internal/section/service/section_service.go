package service

import (
	"context"

	"github.com/cyruzin/meli-frescos/internal/section/domain"
)

type sectionService struct {
	repository domain.SectionRepository
}

func NewSection(sr domain.SectionRepository) domain.SectionService {
	return &sectionService{repository: sr}
}

func (s sectionService) GetAll(ctx context.Context) (*[]domain.Section, error) {
	sections, err := s.repository.GetAll(ctx)
	if err != nil {
		return sections, err
	}

	return sections, nil
}

func (s sectionService) GetByID(ctx context.Context, id int64) (*domain.Section, error) {
	section, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return section, err
	}

	return section, nil
}

func (s sectionService) Store(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	section, err := s.repository.Store(ctx, section)
	if err != nil {
		return section, err
	}

	return section, nil
}

func (s sectionService) Update(ctx context.Context, section *domain.Section) (*domain.Section, error) {
	current, err := s.GetByID(ctx, section.ID)
	if err != nil {
		return section, err
	}

	if section.SectionNumber > 0 {
		current.SectionNumber = section.SectionNumber
	}

	if section.CurrentTemperature > 0 || section.CurrentTemperature < 0 {
		current.CurrentTemperature = section.CurrentTemperature
	}

	if section.MinimumTemperature > 0 || section.MinimumTemperature < 0 {
		current.MinimumTemperature = section.MinimumTemperature
	}

	if section.CurrentCapacity > 0 {
		current.CurrentCapacity = section.CurrentCapacity
	}

	if section.MinimumCapacity > 0 {
		current.MinimumCapacity = section.MinimumCapacity
	}

	if section.MaximumCapacity > 0 {
		current.MaximumCapacity = section.MaximumCapacity
	}

	if section.WarehouseID > 0 {
		current.WarehouseID = section.WarehouseID
	}

	if section.ProductTypeID > 0 {
		current.ProductTypeID = section.ProductTypeID
	}

	section, err = s.repository.Update(ctx, current)
	if err != nil {
		return section, err
	}

	return section, nil
}

func (s sectionService) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
