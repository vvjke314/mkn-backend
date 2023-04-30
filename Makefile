run:
		go run cmd/main.go
swagger:
		swag init --parseDependency --parseInternal -g cmd/main.go
build:
		go build -o bin/main cmd/main.go
dbUp:
		docker-compose up -d
migrate:
		go run cmd/migrate/main.go