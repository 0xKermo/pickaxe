# pickaxe

Indexer of the Starknet AMM pools written in Golang - to be used in [Fibrous](https://fibrous.finance)

<img src="./pickaxe.png" alt="pickaxe girl" width="250px">

*sister of [Shovel](https://github.com/tahos81/shovel) NFT Indexer*

<hr/>

Preperation to build

``` bash
// Create Docker container for Postgres
make postgres

// Creates database in the container
make createdb

// Install migration tool (in macos)
brew install golang-migrate

// Create database tables
make migrateup

// To install the app (preferred)
make install-go

// To build Go code (optional)
make build-go
```

Run the app

``` bash
// Run the app (basic version)
make go

// Run the app (if the app has been installed)
pickaxe

// Run the app (if the code has been built)
./bin/pickaxe

```

Notes:
* You should prepare a initial tokens - pools list for the initial run. The indexer will accept this point as a synced point.

<hr/>

Requirements & Tools
* GoLang (https://go.dev/)
* [dbml](https://dbml-lang.org)
* [docker](https://docker.com/)
* [golang-migrate](https://github.com/golang-migrate/migrate)
* [caigo](https://github.com/dontpanicdao/caigo)
* [caigo-rpcv02](https://github.com/ulerdogan/caigo-rpcv02) (customized rpcv02 of caigo for pickaxe)
* [gocron](https://github.com/go-co-op/gocron)