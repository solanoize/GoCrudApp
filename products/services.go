package products

import "errors"

type ProductService struct {
	productRepository IProductRepository
}

func NewProductService(productRepository IProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (service *ProductService) GetAllProduct() ([]ProductResponseDTO, error) {
	var err error = nil
	var products []Product = []Product{}
	var productResponseDTO ProductResponseDTO = ProductResponseDTO{}

	products, err = service.productRepository.FindAll()
	if err != nil {
		return []ProductResponseDTO{}, errors.New("Gagal mengambil daftar produk")
	}

	return productResponseDTO.toListDTO(products), nil
}

func (service *ProductService) CreateProduct(productRequestDTO ProductRequestDTO) (ProductResponseDTO, error) {
	var err error = nil
	var productResponseDTO ProductResponseDTO = ProductResponseDTO{}
	var product Product = Product{}

	product, err = service.productRepository.Create(productRequestDTO.toModel())
	if err != nil {
		return productResponseDTO, err
	}

	return productResponseDTO.toDTO(product), nil
}

func (service *ProductService) GetProduct(ID string) (ProductResponseDTO, error) {
	var err error = nil
	var product Product = Product{}
	var productResponseDTO ProductResponseDTO = ProductResponseDTO{}

	product, err = service.productRepository.FindByID(ID)
	if err != nil {
		return productResponseDTO, errors.New("Produk tidak ditemukan")
	}

	productResponseDTO.toDTO(product)

	return productResponseDTO, nil
}

func (service *ProductService) UpdateProduct(ID string, requestDTO ProductRequestDTO) (ProductResponseDTO, error) {
	var product Product = Product{}
	var err error
	var responseDTO ProductResponseDTO = ProductResponseDTO{}

	product, err = service.productRepository.FindByID(ID)

	if err != nil {
		return responseDTO, errors.New("Produk tidak ditemukan")
	}

	product.Name = requestDTO.Name
	product.Price = requestDTO.Price
	product.Stock = requestDTO.Stock

	product, err = service.productRepository.Update(ID, product)

	if err != nil {
		return responseDTO, errors.New("Produk tidak dapat diperbarui")
	}

	responseDTO.toDTO(product)

	return responseDTO, nil
}

func (service *ProductService) DeleteProduct(ID string) error {
	var err error = nil

	_, err = service.GetProduct(ID)

	if err != nil {
		return err
	}

	err = service.productRepository.Delete(ID)

	if err != nil {
		return errors.New("Terjadi kesalahan saat menghapus produk")
	}

	return nil
}
