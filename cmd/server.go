package main

import (
	"context"

	"github.com/zakirkun/ehe/internal/config"
	"github.com/zakirkun/ehe/internal/instance"
	"github.com/zakirkun/ehe/internal/services"
)

var configs config.Config
var ctx context.Context

func init() {
	cfg, err := config.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	configs = cfg
	ctx = context.TODO()
}

func main() {
	instance := instance.NewInstance(configs, ctx)
	services := services.NewServices(instance)

	services.WebServer()
}
