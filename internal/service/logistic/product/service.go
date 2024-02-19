package product

import (
	"github.com/AMolchansky/demo_bot/internal/model/logistic"
)

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
	productIndex, err := getProductIndex(productId)

	if err != nil {
		return nil, err
	}

	return &allProducts[productIndex], nil
}

func (s *DummyProductService) Create(product logistic.Product) (uint64, error) {
	allProducts = append(allProducts, product)

	productId := uint64(len(allProducts) - 1)

	return productId, nil
}

func (s *DummyProductService) Update(productID uint64, product logistic.Product) error {
	// find old product from allProducts

	// if not found return error

	// remove old product
	// insert instead of old product new from arg

	return nil
}

func (s *DummyProductService) Remove(productId uint64) (bool, error) {
	productIndex, err := getProductIndex(productId)

	if err != nil {
		return false, err
	}

	//allProducts[productIndex] = nil TODO: добавить в продукты айдишники имитируя работу с бд

	return true, nil
}
