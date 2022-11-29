package services

import (
	"log"

	"golang.org/x/sync/errgroup"
)

func (s IServicesContext) WebServer() {
	var groupRouter errgroup.Group

	groupRouter.Go(func() error {
		return s.instance.WebServerSetup()
	})

	if err := groupRouter.Wait(); err != nil {
		log.Fatalf("Error : %s", err)
	}
}
