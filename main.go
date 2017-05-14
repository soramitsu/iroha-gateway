package main

import "github.com/goadesign/goa"

func main() {
	service := goa.New("Iroha-gateway Server")

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
