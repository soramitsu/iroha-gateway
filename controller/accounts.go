package controller

import (
	"github.com/goadesign/goa"
	"github.com/soramitsu/iroha-gateway/app"
)

// AccountsController implements the accounts resource.
type AccountsController struct {
	*goa.Controller
}

// NewAccountsController creates a accounts controller.
func NewAccountsController(service *goa.Service) *AccountsController {
	return &AccountsController{Controller: service.NewController("AccountsController")}
}

// Add runs the add action.
func (c *AccountsController) Add(ctx *app.AddAccountsContext) error {
	// AccountsController_Add: start_implement

	// Put your logic here

	// AccountsController_Add: end_implement
	return nil
}

// Delete runs the delete action.
func (c *AccountsController) Delete(ctx *app.DeleteAccountsContext) error {
	// AccountsController_Delete: start_implement

	// Put your logic here

	// AccountsController_Delete: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// Get runs the get action.
func (c *AccountsController) Get(ctx *app.GetAccountsContext) error {
	// AccountsController_Get: start_implement

	// Put your logic here

	// AccountsController_Get: end_implement
	res := &app.AnAccount{}
	return ctx.OK(res)
}

// GetAll runs the getAll action.
func (c *AccountsController) GetAll(ctx *app.GetAllAccountsContext) error {
	// AccountsController_GetAll: start_implement

	// Put your logic here

	// AccountsController_GetAll: end_implement
	res := &app.Accounts{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountsController) Update(ctx *app.UpdateAccountsContext) error {
	// AccountsController_Update: start_implement

	// Put your logic here

	// AccountsController_Update: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}
