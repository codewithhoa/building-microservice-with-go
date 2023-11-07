.PHONY:

# install live reload go apps:
air:
	go install github.com/cosmtrek/air@latest
	air init

dev:
	air

call-products:
	curl -X GET http://localhost:9090/products | jq