//+build wireinject

package config

import (
	"context"
	"github.com/google/wire"
	"github.com/mandatorySuicide/golang-code-quality/src/api"
	"github.com/mandatorySuicide/golang-code-quality/src/repository"
	"github.com/mandatorySuicide/golang-code-quality/src/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitProductAPI(ctx context.Context, productCollection mongo.Collection) api.ProductAPI {
	wire.Build(repository.ProvideProductRepository, service.ProvideProductService, api.ProvideProductAPI)
	return api.ProductAPI{}
}
