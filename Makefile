.PHONY:

# install live reload go apps:
air:
	go install github.com/cosmtrek/air@latest
	air init

dev:
	air