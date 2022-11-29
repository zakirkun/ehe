package controller

import internServices "github.com/zakirkun/ehe/internal/services"

type Controller struct {
	services internServices.IServices
}

func InjectServices(s internServices.IServices) Controller {
	return Controller{services: s}
}
