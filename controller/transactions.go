package controller

import (
	"github.com/goadesign/goa"
	"github.com/soramitsu/iroha-gateway/app"
)

// TransactionsController implements the transactions resource.
type TransactionsController struct {
	*goa.Controller
}

// NewTransactionsController creates a transactions controller.
func NewTransactionsController(service *goa.Service) *TransactionsController {
	return &TransactionsController{Controller: service.NewController("TransactionsController")}
}

// GetAll runs the getAll action.
func (c *TransactionsController) GetAll(ctx *app.GetAllTransactionsContext) error {
	// TransactionsController_GetAll: start_implement

	// Put your logic here

	// TransactionsController_GetAll: end_implement
	res := &app.Transactions{}
	return ctx.OK(res)
}
