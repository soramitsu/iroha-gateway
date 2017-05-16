BASE_PACKAGE=github.com/soramitsu/doberman-server

install-glide:
	sh util/glide.sh v0.12.3

deps: install-glide
	glide install

build:
	go build -o bin/iroha-server main.go

goa:
	rm -fr app; goagen app -d github.com/soramitsu/iroha-gateway/design
	goagen controller --pkg='controller' --app-pkg='../app' -o controller -d github.com/soramitsu/iroha-gateway/design
	rm -fr swagger; goagen swagger -d github.com/soramitsu/iroha-gateway/design
	rm -fr client; goagen client -d github.com/soramitsu/iroha-gateway/design