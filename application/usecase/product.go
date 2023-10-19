package usecase

import "github.com/matheuslssilva/fullcycle-grpc-products-challenge/domain/model"

type ProductUseCase struct {
	ProductRepositoryInterface model.ProductRepositoryInterface
}

func (p *ProductUseCase) Create(name string, value float64) (*model.Product, error) {
	product, err := model.NewProduct(name, value)
	if err != nil {
		return nil, err
	}

	productCreated, err := p.ProductRepositoryInterface.Create(product)
	if err != nil {
		return nil, err
	}

	return productCreated, nil
}

func (p *ProductUseCase) FindAll() ([]*model.Product, error) {
	products, err := p.ProductRepositoryInterface.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}
