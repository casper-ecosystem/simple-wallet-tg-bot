# library
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
git clone https://github.com/Simplewallethq/rest-api
cd rest-api/
docker build -t cspr_rest .
docker run -p 8081:8081 cspr_rest    
```

After starting, the swagger interface will be available at:
```
http://localhost:8081/swagger/index.html
```
