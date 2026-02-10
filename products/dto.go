package products

import (
	"errors"
	"gocrud/core"
	"time"
)

type ProductRequestDTO struct {
	Name  string
	Price int64
	Stock int
}

func (dto *ProductRequestDTO) Validate() (core.ValidationErrors, error) {
	var messages core.ValidationErrors = core.ValidationErrors{}

	if dto.Name == "" {
		messages["name"] = "Nama produk tidak boleh kosong"
	}

	if dto.Price <= 0 {
		messages["price"] = "Harga produk tidak boleh kurang dari 0"
	}

	if dto.Stock < 0 {
		messages["stock"] = "Stok produk tidak boleh minus"
	}

	if len(messages) > 0 {
		return messages, errors.New("Validation error")
	}

	return nil, nil
}

func (dto *ProductRequestDTO) toModel() Product {
	return Product{
		Name:  dto.Name,
		Price: dto.Price,
		Stock: dto.Stock,
	}
}

type ProductResponseDTO struct {
	ID        string
	Name      string
	Price     int64
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (dto *ProductResponseDTO) toDTO(model Product) ProductResponseDTO {
	dto.ID = model.ID
	dto.Name = model.Name
	dto.Price = model.Price
	dto.Stock = model.Stock
	dto.CreatedAt = model.CreatedAt
	dto.UpdatedAt = model.UpdatedAt

	return *dto
}

func (dto *ProductResponseDTO) toListDTO(products []Product) []ProductResponseDTO {
	dtos := []ProductResponseDTO{}

	for _, product := range products {
		dtos = append(dtos, dto.toDTO(product))
	}

	return dtos
}
