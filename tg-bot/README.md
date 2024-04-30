# tg-bot
This is a rest interface for library for getting information from the casper blockchain. It can retrieve data such as transaction history, rewards history, and so on.
# Local linter setup

1. Add script to run linter to .git/hooks/pre-commit

```bash
#!/bin/sh

set -e

golangci-lint run ./...
```

2. ```chmod +x .git/hooks/pre-commit```

# Docker

```sh
docker build -t tg-bot .
docker run -d --env TG_TOKEN="TG_TOKEN" --env REST_HOST="https://stage.dspense.com/rest/api/v1/cspr-testnet" --env RPC_NODE="65.108.127.242" CHAIN="casper-test or casper"  --env DB_HOST="PG_DB_HOST" --env DB_PORT="PG_DB_PORT" --env  DB_USER="PG_DB_USER"  --env DB_PASSWORD="PG_DB_PASS" --env SWAP_TOKEN="SIMPLE SWAP TOKEN"  --env DB_NAME="PG_DB_NAME"  tg-bot --env DB_SSLMODE=disable --env PK_SALT="TEST SALT";
```
Before using the bot, you must apply migrations to the postgres database. This project uses the ent orm system and uses "atlas" for migration. The migration files are located in ./ent/migrate/migrations.

To apply migrations run the following command.

```
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "postgres://user:password@localhost:5432/dbname"
```


After starting, the swagger interface will be available at:
```
http://localhost:8081/swagger/index.html
```

# Architecture

The project consists of 2 main parts. botmain and tggateway. Both parts can communicate through any message broker. Protobuff is used as a data serialization format.

## tggateway
Gateway is used for communication with telegram api. It contains a sender and a handler. sender receives messages from botmain and sends them to telegram. handler processes incoming events from telegrams and passes them to botmain.

## botmain

Botmain is responsible for the logical part of this project. it is divided into several parts.

### restclient

Package for working with rest api for interaction with the casper blockchain.

### notificator 

A package for working with notifications for users. Processes events such as incoming transaction, etc.

### taskrecover
Taskrecover is used to recover heavy tasks after a restart.

### userstate
A package for storing the user's state in memory while working with the bot. For example: storing the current address book shift, balance history, etc.

### validators

A module that stores and updates current network validators. Renewal occurs every new era. The following information is stored: Address, activity (true/false), commission, number of delegates.

### invoices

The Invoices feature allows users to create and manage invoices within the Telegram bot framework, leveraging Casper blockchain for transactions. Users can generate invoices with specific parameters including the invoice name, amount in CSPR, comments, and the number of times an invoice can be paid. Each invoice generates a short link directing to the bot for payment, facilitating easy access and user interaction.

Payment options include direct transfer and swap service integration. The system uses a unique memo ID for each invoice to verify transactions. Due to the variability in received amounts when using swap services, an invoice is considered paid if the amount received deviates by no more than 5% from the requested total, accommodating for exchange rate fluctuations.