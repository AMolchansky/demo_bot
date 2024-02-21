package product

import (
	"errors"
	"github.com/AMolchansky/demo_bot/internal/model/logistic"
)

var allProducts = []logistic.Product{
	{Id: 1, Title: "one"},
	{Id: 2, Title: "two"},
	{Id: 3, Title: "three"},
	{Id: 4, Title: "four"},
	{Id: 5, Title: "five"},
	{Id: 6, Title: "six"},
	{Id: 7, Title: "seven"},
}

func getProductById(id uint64) (logistic.Product, error) {
	for _, product := range allProducts {
		if product.Id == id {
			return product, nil
		}
	}

	return logistic.Product{}, errors.New("product not found")
}

func getPaginatedProducts(products []logistic.Product, pageSize int) [][]logistic.Product {
	var result [][]logistic.Product

	for i := 0; i < len(products); i += pageSize {
		limitTo := i + pageSize
		if limitTo > len(products) {
			limitTo = len(products)
		}
		result = append(result, products[i:limitTo])
	}

	return result
}
