# Tools folder

## How can we manage tools version for all member in the team?

> If you want all members of the your team, or CI/CD using the same version of the tools to generate, build.

Example: When you install protocol compiler plugins to generate Go code.
You usually use this command:

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

protoc-gen-go is only a tools for us, it helps us generate go code,
We do not use it in you code. So protoc-gen-go is indirect package, then when
we use command "go mod tidy" => protoc-gen-go will be disappear. If you want to make sure everyone have the same version of protoc-gen-go. This is a way for you

In tools.go file in tools folder from the root of project:

```go
//go:build tools
// +build tools

package tools

import (
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
```

This file is a trick helping you hold protoc-gen-go likes direct package.
But it does not use when building, so we do not worry about this package will be
built with our code base.

## References

- [Go coffeeshop](https://github.com/thangchung/go-coffeeshop/blob/main/tools/tools.go)
- [Link 1](https://marcofranssen.nl/manage-go-tools-via-go-modules)
- [Link 2](https://www.reddit.com/r/golang/comments/10rlp31/toolsgo_pattern_still_valid_today_i_want_to/)
