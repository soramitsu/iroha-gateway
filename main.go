package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/soramitsu/iroha-gateway/app"
	. "github.com/soramitsu/iroha-gateway/controller"
)

func main() {
	service := goa.New("Iroha-gateway Server")

	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	c1 := NewCurrencyController(service)
	app.MountCurrencyController(service, c1)

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
