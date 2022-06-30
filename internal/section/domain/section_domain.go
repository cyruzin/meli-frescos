package domain


import (
	"context"
)

// Section model
//
// id: id único da seção
// current_temperature: Uma temperatura atual
// minimum_temperature: Uma temperatura mínima
// current_capacity: Uma capacidade atual, dependendo dos lotes em
// um determinado momento.
// minimum_capacity: uma capacidade mínima
// maximum_capacity: Uma capacidade máxima, dependendo do
// tamanho do setor
// warehouse_id: um armazém associado (wareHouse)
// product_type_id: um tipo de produto associado (ProductType)
type Section struct {
	ID                 int64 `json:"id"`
	SectionNumber      int64 `json:"section_number"`
	CurrentTemperature int16 `json:"current_temperature"`
	MinimumTemperature int16 `json:"minimum_temperature"`
	CurrentCapacity    int64 `json:"current_capacity"`
	MinimumCapacity    int64 `json:"minimum_capacity"`
	MaximumCapacity    int64 `json:"maximum_capacity"`
	WarehouseID        int64 `json:"warehouse_id"`
	ProductTypeID      int64 `json:"product_type_id"`
}

type SectionRepository interface {
	GetAll(ctx context.Context) (*[]Section, error)
	GetByID(ctx context.Context, id int64) (*Section, error)
	Store(ctx context.Context, section *Section) (*Section, error)
	Update(ctx context.Context, section *Section) (*Section, error)
	Delete(ctx context.Context, id int64) error
}

type SectionService interface {
	GetAll(ctx context.Context) (*[]Section, error)
	GetByID(ctx context.Context, id int64) (*Section, error)
	Store(ctx context.Context, section *Section) (*Section, error)
	Update(ctx context.Context, section *Section) (*Section, error)
	Delete(ctx context.Context, id int64) error
}
