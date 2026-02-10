package products

import (
	"gorm.io/gorm"
)

type IProductRepository interface {
	FindAll() ([]Product, error)
	FindByID(ID string) (Product, error)
	Create(product Product) (Product, error)
	Update(ID string, product Product) (Product, error)
	Delete(ID string) error
}

type ProductRepository struct {
	db *gorm.DB
}

func (p *ProductRepository) Create(product Product) (Product, error) {
	var err error = nil
	err = p.db.Create(&product).Error

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (p *ProductRepository) Delete(ID string) error {
	var err error = nil

	err = p.db.Delete(&Product{}, "id = ?", ID).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) FindAll() ([]Product, error) {
	var err error = nil
	var products []Product = []Product{}

	err = p.db.Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (p *ProductRepository) FindByID(ID string) (Product, error) {
	var err error = nil
	var product Product = Product{}

	err = p.db.First(&product, "id = ?", ID).Error

	if err != nil {
		return product, err
	}

	return product, nil

}

func (p *ProductRepository) Update(ID string, product Product) (Product, error) {
	var err error = nil

	err = p.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func ProductRepositoryImpl(db *gorm.DB) IProductRepository {
	return &ProductRepository{db: db}
}
