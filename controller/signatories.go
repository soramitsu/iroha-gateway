package controller

import (
	"github.com/goadesign/goa"
	"github.com/soramitsu/iroha-gateway/app"
)

// SignatoriesController implements the signatories resource.
type SignatoriesController struct {
	*goa.Controller
}

// NewSignatoriesController creates a signatories controller.
func NewSignatoriesController(service *goa.Service) *SignatoriesController {
	return &SignatoriesController{Controller: service.NewController("SignatoriesController")}
}

// Add runs the add action.
func (c *SignatoriesController) Add(ctx *app.AddSignatoriesContext) error {
	// SignatoriesController_Add: start_implement

	// Put your logic here

	// SignatoriesController_Add: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// Delete runs the delete action.
func (c *SignatoriesController) Delete(ctx *app.DeleteSignatoriesContext) error {
	// SignatoriesController_Delete: start_implement

	// Put your logic here

	// SignatoriesController_Delete: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// GetAll runs the getAll action.
func (c *SignatoriesController) GetAll(ctx *app.GetAllSignatoriesContext) error {
	// SignatoriesController_GetAll: start_implement

	// Put your logic here

	// SignatoriesController_GetAll: end_implement
	res := &app.Signatories{}
	return ctx.OK(res)
}
