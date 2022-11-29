package services

import "github.com/zakirkun/ehe/internal/instance"

type IServicesContext struct {
	instance instance.IAppContext
}

type iServices interface {
	WebServer()
}

func NewServices(i instance.IAppContext) iServices {
	return &IServicesContext{instance: i}
}
