run:
		go run cmd/main.go
swagger:
		export PATH=$(go env GOPATH)/bin:$(PATH)
		swag init --parseDependency --parseInternal -g cmd/main.go
build:
		go build -o bin/main cmd/main.go

