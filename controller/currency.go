package controller

import (
	"github.com/goadesign/goa"
	"github.com/soramitsu/iroha-gateway/app"
)

// CurrencyController implements the currency resource.
type CurrencyController struct {
	*goa.Controller
}

// NewCurrencyController creates a currency controller.
func NewCurrencyController(service *goa.Service) *CurrencyController {
	return &CurrencyController{Controller: service.NewController("CurrencyController")}
}

// Add runs the add action.
func (c *CurrencyController) Add(ctx *app.AddCurrencyContext) error {
	// CurrencyController_Add: start_implement

	// Put your logic here

	// CurrencyController_Add: end_implement
	return nil
}

// AddValue runs the addValue action.
func (c *CurrencyController) AddValue(ctx *app.AddValueCurrencyContext) error {
	// CurrencyController_AddValue: start_implement

	// Put your logic here

	// CurrencyController_AddValue: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// Delete runs the delete action.
func (c *CurrencyController) Delete(ctx *app.DeleteCurrencyContext) error {
	// CurrencyController_Delete: start_implement

	// Put your logic here

	// CurrencyController_Delete: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// GetAll runs the getAll action.
func (c *CurrencyController) GetAll(ctx *app.GetAllCurrencyContext) error {
	// CurrencyController_GetAll: start_implement

	// Put your logic here

	// CurrencyController_GetAll: end_implement
	res := &app.Currency{}
	return ctx.OK(res)
}

// SubtractValue runs the subtractValue action.
func (c *CurrencyController) SubtractValue(ctx *app.SubtractValueCurrencyContext) error {
	// CurrencyController_SubtractValue: start_implement

	// Put your logic here

	// CurrencyController_SubtractValue: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// Transfer runs the transfer action.
func (c *CurrencyController) Transfer(ctx *app.TransferCurrencyContext) error {
	// CurrencyController_Transfer: start_implement

	// Put your logic here

	// CurrencyController_Transfer: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *CurrencyController) Update(ctx *app.UpdateCurrencyContext) error {
	// CurrencyController_Update: start_implement

	// Put your logic here

	// CurrencyController_Update: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}
