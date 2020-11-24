package domain

type Product struct {
	BaseEntity `bson:",inline"`
	Code       string `bson:"code"`
	Price      int    `bson:"price"`
}

func NewProduct(baseEntity *BaseEntity, code string, price int) *Product {
	return &Product{BaseEntity: *baseEntity, Code: code, Price: price}
}
