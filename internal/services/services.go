package services

import (
	"github.com/zakirkun/ehe/internal/instance"
	"go.mongodb.org/mongo-driver/mongo"
)

type IServicesContext struct {
	instance instance.IAppContext
}

type IServices interface {
	WebServer()
	OpenDB() *mongo.Client
}

func NewServices(i instance.IAppContext) IServices {
	return &IServicesContext{instance: i}
}
