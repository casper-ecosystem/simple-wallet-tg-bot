This repository contains 3 sub-repositories

library - library for working with blockchain casper
rest - rest api for working with the Casper blockchain
tg-bot - telegram wallet

To run telegram wallet you need to run rest api

#REST
## Docker

```sh
cd rest
docker build --build-arg GITHUB_ACCESS_TOKEN="your_github_token"  -t cspr_rest .
docker run -p 8081:8081 cspr_rest    
```

After starting, the swagger interface will be available at:
```
http://localhost:8081/swagger/index.html
```


# run telegram bot


```sh
cd tgbot
docker build -t tg-bot .
docker run  --env TG_TOKEN="bot token" --env REST_HOST="127.0.0.1/api/v1/cspr-testnet" --env RPC_NODE="213.239.207.87" --env CHAIN="casper-test"  --env DB_HOST="dbhost" --env DB_PORT="54320" --env  DB_USER="user"  --env DB_PASSWORD="password" --env SWAP_TOKEN="SIMPLE SWAP TOKEN"  --env DB_NAME="tgbot"  --env DB_SSLMODE=disable --env PK_SALT="TEST SALT" tg-bot```