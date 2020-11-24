package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mandatorySuicide/golang-code-quality/src/service"
)

type ProductAPI struct {
	ProductService service.ProductService
}

func ProvideProductAPI(p service.ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	products, err := p.ProductService.FindAll()
	if err != nil {
		c.JSON(err.Status, err)
	} else {
		c.JSON(200, products)
	}
}
