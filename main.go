package main

import "github.com/goadesign/goa"

func main() {
	service := goa.New("Iroha-gateway Server")

	if err := service.ListenAndServe(":9090"); err != nil {
		service.LogError("startup", "err", err)
	}
}
