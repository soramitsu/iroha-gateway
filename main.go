package main

import (
	"github.com/goadesign/goa"
	"github.com/soramitsu/iroha-gateway/controller"
	"github.com/soramitsu/iroha-gateway/app"
)

func main() {
	service := goa.New("Iroha-gateway Server")

	accounts := controller.NewAccountController(service)
	app.MountAccountController(service, accounts)

	currency := controller.NewCurrencyController(service)
	app.MountCurrencyController(service, currency)

	quorum := controller.NewQuorumController(service)
	app.MountQuorumController(service, quorum)

	signatories := controller.NewSignatoriesController(service)
	app.MountSignatoriesController(service, signatories)

	transactions := controller.NewTransactionsController(service)
	app.MountTransactionsController(service, transactions)

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
