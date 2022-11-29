package services

type IServicesContext struct {
	userServices usersServiceContext
	authServices authServicesContext
}

func NewServices() IServicesContext {
	return IServicesContext{}
}
