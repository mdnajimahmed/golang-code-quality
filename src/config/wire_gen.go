// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package config

import (
	"context"
	"github.com/mandatorySuicide/golang-code-quality/src/api"
	"github.com/mandatorySuicide/golang-code-quality/src/repository"
	"github.com/mandatorySuicide/golang-code-quality/src/service"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitProductAPI(ctx context.Context, productCollection mongo.Collection) api.ProductAPI {
	productRepository := repository.ProvideProductRepository(ctx, productCollection)
	productService := service.ProvideProductService(productRepository)
	productAPI := api.ProvideProductAPI(productService)
	return productAPI
}
