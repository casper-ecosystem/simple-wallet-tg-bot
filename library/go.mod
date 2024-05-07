module github.com/Simplewallethq/source-code/library

go 1.20

//replace github.com/Simplewallethq/rest-api => ../rest-api

require (
	github.com/google/go-cmp v0.6.0
	github.com/make-software/casper-go-sdk v1.5.1
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace github.com/casper-ecosystem/casper-golang-sdk => ../casper-golang-sdk

// replace github.com/casper-ecosystem/casper-golang-sdk v0.0.0-20220131100450-718690b51142 => github.com/Simplewallethq/casper-golang-sdk v0.0.0-20230530122749-7a3a6197b733
