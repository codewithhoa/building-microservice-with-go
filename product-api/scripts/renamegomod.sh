#!/bin/bash

set -euo pipefail

OLD_MODULE_NAME="github.com/codewithhoa/building-microservice-with-go"
NEW_MODULE_NAME="github.com/codewithhoa/building-microservice-with-go/product-api"

go mod edit -module $NEW_MODULE_NAME
