package controller

import (
	"github.com/goadesign/goa"
	"github.com/soramitsu/iroha-gateway/app"
)

// QuorumController implements the quorum resource.
type QuorumController struct {
	*goa.Controller
}

// NewQuorumController creates a quorum controller.
func NewQuorumController(service *goa.Service) *QuorumController {
	return &QuorumController{Controller: service.NewController("QuorumController")}
}

// Update runs the update action.
func (c *QuorumController) Update(ctx *app.UpdateQuorumContext) error {
	// QuorumController_Update: start_implement

	// Put your logic here

	// QuorumController_Update: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}
