package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (service *ProductService) GetAll() ([]models.Product, error) {
	return service.repo.GetAll()
}

func (service *ProductService) Create(data *models.Product) error {
	return service.repo.Create(data)
}

func (service *ProductService) GetByID(id int) (*models.Product, error){
	return service.repo.GetByID(id)
}

func (service *ProductService) Update(product *models.Product) error {
	return service.repo.Update(product)
}

func (service *ProductService) Delete(id int) error {
	return service.repo.Delete(id)
}
