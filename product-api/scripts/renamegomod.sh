# This script allow change go module name
# and automatically change old path from import statement in all .go file
# ref:
# - https://gist.github.com/iamucil/7578dc7df7d72e1d78c8f5543db3fbcc
# - https://www.cyberciti.biz/faq/how-to-use-sed-to-find-and-replace-text-in-files-in-linux-unix-shell/
# - https://www.geeksforgeeks.org/sed-command-in-linux-unix-with-examples/
# - https://www.geeksforgeeks.org/find-command-in-linux-with-examples/
#!/bin/bash

set -euo pipefail

OLD_MODULE_NAME="github.com/codewithhoa/building-microservice-with-go"
NEW_MODULE_NAME="github.com/codewithhoa/building-microservice-with-go/product-api"

go mod edit -module $NEW_MODULE_NAME

find . -type f -name '*.go' \
  -exec sed -i '' "s+$OLD_MODULE_NAME+$NEW_MODULE_NAME+g" {} \;
