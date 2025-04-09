package usecase

import (
	"inventory-service/domain"
	"inventory-service/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService() *Service {
	return &Service{repository: repository.NewRepository()}
}

func (s *Service) GetAllProducts() ([]domain.Product, error) {
	return s.repository.GetAllProducts()
}

func (s *Service) GetProductByID(id uint) (*domain.Product, error) {
	return s.repository.GetProductByID(id)
}

func (s *Service) CreateProduct(product domain.Product) (*domain.Product, error) {
	return s.repository.CreateProduct(product)
}

func (s *Service) UpdateProduct(product domain.Product) (*domain.Product, error) {
	return s.repository.UpdateProduct(product)
}

func (s *Service) DeleteProduct(id uint) error {
	return s.repository.DeleteProduct(id)
}
