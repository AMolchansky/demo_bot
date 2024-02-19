package product

import (
	"errors"
	"github.com/AMolchansky/demo_bot/internal/model/logistic"
)

var allProducts = []logistic.Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

func getProductIndex(productId uint64) (int, error) {
	for index := range allProducts {
		if uint64(index) == productId {
			return index, nil
		}
	}

	return -1, errors.New("product not found")
}
