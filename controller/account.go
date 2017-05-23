package controller

import (
	"github.com/goadesign/goa"
	"github.com/soramitsu/iroha-gateway/app"
)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service) *AccountController {
	return &AccountController{Controller: service.NewController("AccountController")}
}

// Add runs the add action.
func (c *AccountController) Add(ctx *app.AddAccountContext) error {
	// AccountController_Add: start_implement

	// Put your logic here

	// AccountController_Add: end_implement
	return nil
}

// Delete runs the delete action.
func (c *AccountController) DeleteByUsername(ctx *app.DeleteByUsernameAccountContext) error {
	// AccountController_Delete: start_implement

	// Put your logic here

	// AccountController_Delete: end_implement
	res := &app.Messageresponse{}
	return ctx.OK(res)
}

func (c *AccountController) DeleteByUUID(ctx *app.DeleteByUUIDAccountContext) error {
	// AccountController_Delete: start_implement

	// Put your logic here

	// AccountController_Delete: end_implement
	res := &app.Messageresponse{}
	return ctx.OK(res)
}

func (c *AccountController) DeleteByUsernameFromDefaultDomain(ctx *app.DeleteByUsernameFromDefaultDomainAccountContext) error {
	// AccountController_Delete: start_implement

	// Put your logic here

	// AccountController_Delete: end_implement
	res := &app.Messageresponse{}
	return ctx.OK(res)
}

// GetAll runs the getAll action.
func (c *AccountController) GetAll(ctx *app.GetAllAccountContext) error {
	// AccountController_GetAll: start_implement

	// Put your logic here

	// AccountController_GetAll: end_implement
	res := &app.Accountsresponse{}
	return ctx.OK(res)
}

// GetByUUID runs the getByUUID action.
func (c *AccountController) GetByUUID(ctx *app.GetByUUIDAccountContext) error {
	// AccountController_GetByUUID: start_implement

	// Put your logic here

	// AccountController_GetByUUID: end_implement
	res := &app.Accountresponse{}
	return ctx.OK(res)
}

// GetByUsername runs the getByUsername action.
func (c *AccountController) GetByUsername(ctx *app.GetByUsernameAccountContext) error {
	// AccountController_GetByUsername: start_implement

	// Put your logic here

	// AccountController_GetByUsername: end_implement
	res := &app.Accountresponse{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountController) GetByUsernameFromDefaultDomain(ctx *app.GetByUsernameFromDefaultDomainAccountContext) error {
	// AccountController_Update: start_implement

	// Put your logic here

	// AccountController_Get: end_implement
	res := &app.Accountresponse{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountController) UpdateByUsername(ctx *app.UpdateByUsernameAccountContext) error {
	// AccountController_Update: start_implement

	// Put your logic here

	// AccountController_Update: end_implement
	res := &app.Messageresponse{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountController) UpdateByUUID(ctx *app.UpdateByUUIDAccountContext) error {
	// AccountController_Update: start_implement

	// Put your logic here

	// AccountController_Update: end_implement
	res := &app.Messageresponse{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountController) UpdateByUsernameFromDefaultDomain(ctx *app.UpdateByUsernameFromDefaultDomainAccountContext) error {
	// AccountController_Update: start_implement

	// Put your logic here

	// AccountController_Update: end_implement
	res := &app.Messageresponse{}
	return ctx.OK(res)
}

