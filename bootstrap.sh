rm -fr app; goagen app -d github.com/soramitsu/doberman-server/design
goagen controller --pkg='controller' --app-pkg='../app' -o controller -d github.com/soramitsu/doberman-server/design
rm -fr swagger; goagen swagger -d github.com/soramitsu/doberman-server/design
rm -fr client; goagen client -d github.com/soramitsu/doberman-server/design
