package instance

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (i *IAppContext) MongoSetup() *mongo.Client {

	// create connection to mongo
	connection := options.Client().ApplyURI(i.Cfg.DBUri)

	// connect to mongo
	mongoClient, err := mongo.Connect(i.Ctx, connection)
	if err != nil {
		panic(err)
	}

	if err := mongoClient.Ping(i.Ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return mongoClient
}
