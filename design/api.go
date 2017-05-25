package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Iroha-Gateway Server", func() {
	Title("The Iroha API")
	Host("localhost:8080")
	Scheme("http")
})
