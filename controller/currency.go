package controller

import (
	"strconv"

	"github.com/goadesign/goa"
	"github.com/soramitsu/iroha-gateway/app"
	"github.com/soramitsu/iroha-gateway/command"
	"github.com/soramitsu/iroha-gateway/model"
	"github.com/soramitsu/iroha-gateway/service"
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
	p := ctx.Payload
	timestamp, err := strconv.ParseInt(p.Timestamp, 10, 64)
	if err != nil {
		res := &app.Message{Code: 0, Message: err.Error()} //TODO: パース失敗時とかの Code は何？
		return ctx.InternalServerError(res)
	}
	signatures := []model.Signatures{} // TODO: JSONは1つしかシグネチャを受けとらない -> 配列？？
	tx := &model.Transaction{
		Command: &command.Currency{
			CurrencyName: p.Currency.Name,
			DomainName:   p.Currency.DomainName,
			LedgerName:   p.Currency.LedgerName,
			Description:  *p.Currency.Description,
			Amount:       "", // TODO: Payloadにない項目
			Precision:    1,  // TODO: Payloadにない項目
		},
		Creator:    p.CreatorPubkey,
		Hash:       "",         // TODO: Hashとは？？？
		Signatures: signatures, // TODO: JSONにはひとつしかない -> 配列？？ 現状空の配列
		Timestamp:  timestamp,
	}

	service.CreateCurrency(tx)

	// CurrencyController_Add: end_implement
	res := &app.Message{} // TODO: 成功時に何返す？ HTTP statusみたいに200？？？
	return ctx.Created(res)
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
	res := &app.AnCurrency{}
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
