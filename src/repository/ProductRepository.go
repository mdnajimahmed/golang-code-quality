package repository

import (
	"context"
	"github.com/mandatorySuicide/golang-code-quality/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection mongo.Collection
	context    context.Context
}

func ProvideProductRepository(ctx context.Context, mongoCollection mongo.Collection) ProductRepository {
	return ProductRepository{
		context:    ctx,
		collection: mongoCollection,
	}
}

func (p *ProductRepository) Save(product *domain.Product) (*mongo.InsertOneResult, error) {
	return p.collection.InsertOne(p.context, product)
}

func (p *ProductRepository) FindAll() ([]*domain.Product, error) {
	cursor, err := p.collection.Find(p.context, bson.D{})
	defer cursor.Close(p.context)
	if err != nil {
		return nil, err
	}
	list := make([]*domain.Product, 0)
	for cursor.Next(p.context) {
		// Declare a result BSON object
		var result domain.Product
		decodeErr := cursor.Decode(&result)
		if decodeErr != nil {
			return nil, decodeErr
		}
		list = append(list, &result)
	}
	return list,nil

}
