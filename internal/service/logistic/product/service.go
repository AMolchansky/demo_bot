package product

import (
	"errors"
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
	paginatedProducts := getPaginatedProducts(allProducts, int(limit))

	for page, products := range paginatedProducts {
		if uint64(page) == cursor {
			return products, nil
		}
	}

	return []logistic.Product{}, errors.New("invalid page")
}

func (s *DummyProductService) Describe(productId uint64) (*logistic.Product, error) {
	product, err := getProductById(productId)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *DummyProductService) Create(product logistic.Product) (uint64, error) {
	var lastProductId uint64
	if len(allProducts) > 1 {
		lastProductId = allProducts[len(allProducts)-1].Id
	} else {
		lastProductId = 1
	}

	product.Id = lastProductId + 1

	allProducts = append(allProducts, product)

	return product.Id, nil
}

func (s *DummyProductService) Update(productId uint64, product logistic.Product) error {
	_, err := getProductById(productId)

	if err != nil {
		return err
	}

	for i, existProduct := range allProducts {
		if existProduct.Id == productId {
			allProducts[i] = product
		}
	}

	return nil
}

func (s *DummyProductService) Remove(productId uint64) (bool, error) {
	_, err := getProductById(productId)

	if err != nil {
		return false, err
	}

	for i, product := range allProducts {
		if product.Id == productId {
			allProducts = append(allProducts[:i], allProducts[i+1:]...)
		}
	}

	return true, nil
}
