run:
		go run cmd/mkn/main.go
swagger:
		swag init -g cmd/mkn/main.go
build:
		go build -o bin/main cmd/mkn/main.go
dbUp:
		docker-compose up -d
migrate:
		go run cmd/migrate/main.go