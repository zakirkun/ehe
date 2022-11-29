package services

import (
	"github.com/zakirkun/ehe/internal/instance"
	"go.mongodb.org/mongo-driver/mongo"
)

type IServicesContext struct {
	instance instance.IAppContext
}

type iServices interface {
	WebServer()
	OpenDB() *mongo.Client
}

func NewServices(i instance.IAppContext) iServices {
	return &IServicesContext{instance: i}
}
