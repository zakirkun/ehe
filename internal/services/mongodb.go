package services

import "go.mongodb.org/mongo-driver/mongo"

func (i IServicesContext) OpenDB() *mongo.Client {
	return i.instance.MongoSetup()
}
