SHELL=/bin/bash
BASE_PACKAGE=github.com/soramitsu/doberman-server

run-iroha:
	docker-compose -f iroha-docker-compose.yaml down
	docker-compose -f iroha-docker-compose.yaml up -d && docker-compose -f iroha-docker-compose.yaml logs -f

install-glide:
	sh util/glide.sh v0.12.3

deps: install-glide
	glide install

test: 
	go test -v -race ./service

build:
	go build -o bin/iroha-server main.go

goa:
	rm -fr app; goagen app -d github.com/soramitsu/iroha-gateway/design
	goagen controller --pkg='controller' --app-pkg='../app' -o controller -d github.com/soramitsu/iroha-gateway/design
	rm -fr swagger; goagen swagger -d github.com/soramitsu/iroha-gateway/design
	rm -fr client; goagen client -d github.com/soramitsu/iroha-gateway/design
