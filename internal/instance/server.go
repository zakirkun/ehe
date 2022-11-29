package instance

import (
	"net/http"

	"github.com/zakirkun/ehe/app/router"
)

func (i *IAppContext) WebServerSetup() error {
	server := http.Server{
		Addr:    ":" + i.Cfg.Port,
		Handler: router.InitRouters(i.Cfg.Debug),
	}

	return server.ListenAndServe()
}
