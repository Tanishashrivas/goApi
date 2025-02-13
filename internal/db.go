package internal

import (
	"context"
	"fmt"

	utils "github.com/tanishashrivas/goApi/pkg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://muskan:muskan01@cluster0.6dovp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "Udemy"
const colName = "courselist"

var Collection *mongo.Collection

func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionstring)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	utils.CheckNilError(err)
	fmt.Println("Mongodb connection successfull!")

	Collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready", Collection, client)
}
