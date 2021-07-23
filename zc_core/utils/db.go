import "go.mongodb.org/mongo-driver/mongo"


clientOptions := options.Client().
    ApplyURI("mongodb+srv://zuri:<password>@cluster0.hepte.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}
