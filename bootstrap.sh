#!/bin/bash

rm -fr app; goagen app -d github.com/soramitsu/iroha-gateway/design
goagen controller --pkg='controller' --app-pkg='../app' -o controller -d github.com/soramitsu/iroha-gateway/design
rm -fr swagger; goagen swagger -d github.com/soramitsu/iroha-gateway/design
rm -fr client; goagen client -d github.com/soramitsu/iroha-gateway/design
