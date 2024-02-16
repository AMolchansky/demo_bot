package product

import "github.com/AMolchansky/demo_bot/internal/model/logistic"

type ProductService interface {
	Describe(productID uint64) (*logistic.Product, error)
	List(cursor uint64, limit uint64) ([]logistic.Product, error)
	Create(logistic.Product) (uint64, error)
	Update(productID uint64, Product logistic.Product) error
	Remove(productID uint64) (bool, error)
}

type DummyProductService struct {
}

func NewDummyProductService() *DummyProductService {
	return &DummyProductService{}
}

func (s *DummyProductService) List() []Product {
	return allProducts
}

func (s *DummyProductService) Get(idx int) (*Product, error) {
	//todo homework: проверить границы обработать ошибку

	return &allProducts[idx], nil
}
