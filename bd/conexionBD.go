package bd

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ClientMongo obj de conexion */
var ClientMongo *mongo.Database

/*ConectarBD conexion BD*/
func ConectarBD() {
	if ClientMongo == nil {
		uri, database := getUri()
		/*log.Fatal(uri)*/
		client, err := mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err.Error())
		}
		ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		ClientMongo = client.Database(database)
		log.Println("Conexion exitosa a la BD")
	}
}

/*getUri obtiene url*/
func getUri() (string, string) {
	uri, database := os.Getenv("MONGO_URI"), os.Getenv("DATABASE")
	if uri == "" {
		/*uri = "mongodb+srv://sa:Nicolasgarcia03+@cluster0.nllgqfu.mongodb.net/proyecto-go?retryWrites=true&w=majority"*/
		uri = "mongodb://localhost:27017"
	}
	if database == "" {
		database = "proyecto-go"
	}
	return uri, database
}

/*GetCollection obtiene coleccion de mongo*/
func GetCollection(colecction string) *mongo.Collection {
	if ClientMongo == nil {
		return nil
	} else {
		return ClientMongo.Collection(colecction)
	}
}
