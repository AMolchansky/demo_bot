package product

import "github.com/AMolchansky/demo_bot/internal/model/logistic"

type ProductService interface {
	Describe(productID uint64) (*logistic.Product, error)
	List(cursor uint64, limit uint64) ([]logistic.Product, error)
	Create(product logistic.Product) (uint64, error)
	Update(productID uint64, product logistic.Product) error
	Remove(productID uint64) (bool, error)
}

type DummyProductService struct {
}

func NewDummyProductService() *DummyProductService {
	return &DummyProductService{}
}

func (s *DummyProductService) List(cursor uint64, limit uint64) ([]logistic.Product, error) {
	// slice allProducts using cursor limit

	// if limit more than maximum size of products => good
	// if cursor more than limit return error

	return allProducts, nil
}

func (s *DummyProductService) Describe(productId uint64) (*logistic.Product, error) {
	// find products in allProducts

	// if not found return error

	return &allProducts[productId], nil
}

func (s *DummyProductService) Create(product logistic.Product) (uint64, error) {
	// add to all products product from arg

	return 1, nil
}

func (s *DummyProductService) Update(productID uint64, product logistic.Product) error {
	// find old product from allProducts

	// if not found return error

	// remove old product
	// insert instead of old product new from arg

	return nil
}

func (s *DummyProductService) Remove(productId uint64) (bool, error) {
	// find product in allProducts

	// if not found return error

	// remove from allProducts

	return true, nil
}
