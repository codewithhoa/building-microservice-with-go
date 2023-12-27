.PHONY:

# Install tool
tool:
	brew install golangci-lint \
	brew install curl \
	brew install jq

# Install live reload go apps:
air:
	go install github.com/cosmtrek/air@latest \
	air init

lint:
	golangci-lint run

dev:
	air

# Environment variable
SERVER_NAME    = products-api
SERVER_ADDRESS = :9090
SERVER_VERSION = v3.0.0
SERVER_GRACEFULTIMEOUT = 30s
LOG_LEVEL      = Debug

prod:
	go run ./cmd/main.go \
	-sv_name=$(SERVER_NAME) \
	-sv_address=$(SERVER_ADDRESS) \
	-sv_version=$(SERVER_VERSION) \
	-graceful-timeout=$(SERVER_GRACEFULTIMEOUT) \
	-log_level=$(LOG_LEVEL)
		
# Testting api 
END_POINT   = http://localhost:9090/api/v1/products
PRODUCT_ID  = 2
# DATA 			  = '{"name": "tea", "description": "awesome tea for your new day!", "price": 0, "sku": "xxx-xxx-aaa"}'
DATA 			  = '{"name": "aaaa", "description": "awesome tea for your new day!", "price": 0, "sku": "1"}'
DATA_UPDATE = '{"name": "ice cream", "description": "awesome strawberry ice-cream for your new day!", "price": 5, "sku": "icec202"}'

get-products:
	curl -X GET $(END_POINT) | jq

get-product-byID:
	curl -X GET $(END_POINT)/$(id) | jq

post-product:
	curl -v -X POST $(END_POINT) -H "Content-Type: application/json" -d $(DATA) 

put-product:
	curl -v -X PUT $(END_POINT)/$(id) -H "Content-Type: application/json" -d $(DATA) | jq 

del-product:
	curl -v -X DELETE $(END_POINT)/$(id) | jq 


# Swagger
swagger-gen:
	swagger generate spec --input ./api/tags.json -o ./api/swagger.yml --scan-models

swagger-ui: swagger-gen
	swagger serve --port=9999 --flavor=swagger ./api/swagger.yml