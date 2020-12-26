package main
//
//import (
//	"context"
//	"github.com/gin-gonic/gin"
//	"github.com/mandatorySuicide/golang-code-quality/src/config"
//	"github.com/mandatorySuicide/golang-code-quality/src/langPractice"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"os"
//)
//
//type S struct {
//}
//type I interface {
//}
//
//func GetClient(ctx context.Context) *mongo.Client {
//	uri := os.Getenv("MONGO_CONNECTION_STRING")
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
//	if err != nil {
//		panic("Could not connect to database.Try after some time.")
//	}
//	return client
//}
//
//func GetDataBase(ctx context.Context) *mongo.Database {
//	return GetClient(ctx).Database(os.Getenv("DATABASE"))
//}
//
//func GetCollection(ctx context.Context, collection string) *mongo.Collection {
//	return GetDataBase(ctx).Collection(collection)
//}
//
//func main() {
//	langPractice.TestTag()
//	//langPractice.TestSort()
//	langPractice.ConcMain()
//	//oldMain()
//
//}
//
//func oldMain() {
//	ctx := context.Background()
//	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	//defer cancel()
//	productCollection := GetCollection(ctx, "product")
//	////fmt.Println("MAIN INIT")
//	//userId := primitive.NewObjectID()
//	//prod := domain.NewProduct(domain.NewBaseEntity(userId), "ABC123", 100)
//	////insert := InsertIn(prodCol)
//	////saveResult,err := insert(prod)
//	//prodRepo := repository.ProvideProductRepository(ctx,productCollection)
//	//saveResult, err := prodRepo.Save(prod)
//	//fmt.Print(saveResult)
//	//fmt.Print(err)
//	//app.SetupServer().Run()
//
//	productAPI := config.InitProductAPI(ctx, *productCollection)
//	r := gin.Default()
//
//	r.GET("/products", productAPI.FindAll)
//	err := r.Run()
//	if err != nil {
//		panic(err)
//	}
//}
//
