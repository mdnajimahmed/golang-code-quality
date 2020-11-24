package service

import (
	"fmt"
	"github.com/mandatorySuicide/golang-code-quality/src/domain"
	"github.com/mandatorySuicide/golang-code-quality/src/repository"
	"github.com/mandatorySuicide/golang-code-quality/src/utility"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func ProvideProductService(p repository.ProductRepository) ProductService {
	return ProductService{
		ProductRepository: p,
	}
}

func (p *ProductService) Save(product *domain.Product) *utility.RestError {
	insertResult, error := p.ProductRepository.Save(product)
	if error != nil {
		return utility.NewInternalServerError("Insertion failed", &error)
	}
	fmt.Println(insertResult)
	return nil
}

func (p *ProductService) FindAll() ([]*domain.Product, *utility.RestError) {
	products, error := p.ProductRepository.FindAll()
	if error != nil {
		return nil, utility.NewInternalServerError("Search failed", &error)
	}
	return products, nil
}
