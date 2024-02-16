package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	//todo homework: проверить границы обработать ошибку

	return &allProducts[idx], nil
}
