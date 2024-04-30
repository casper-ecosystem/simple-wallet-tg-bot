# library
This is a library for getting information from the casper blockchain. It can retrieve data such as transaction history, rewards history, and so on.
# Local linter setup

1. Add script to run linter to .git/hooks/pre-commit

```bash
#!/bin/sh

set -e

golangci-lint run ./...
```

2. ```chmod +x .git/hooks/pre-commit```