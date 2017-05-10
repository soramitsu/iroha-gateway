BASE_PACKAGE=github.com/soramitsu/doberman-server

install-glide:
	sh util/glide.sh v0.12.3

deps: install-glide
	glide install

build:
	go build -o bin/iroha-server main.go

