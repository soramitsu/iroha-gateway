package service

// TODO:
// client <- gateway: app.Transaction
// gateway controller <- service: service.TransactionResponse?
type TransactionResponse struct {
	Code    byte
	Message string
}
