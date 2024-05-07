This repository contains 3 sub-repositories

library - library for working with blockchain casper
rest - rest api for working with the Casper blockchain
tg-bot - telegram wallet

To run telegram wallet you need to run rest api
# Docker

```sh
cd rest
docker build --build-arg GITHUB_ACCESS_TOKEN="your_github_token"  -t cspr_rest .
docker run -p 8081:8081 cspr_rest    
```

