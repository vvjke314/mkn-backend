FROM golang:latest

RUN go version 
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o ./bin/main ./cmd/mkn/main.go

CMD ["./bin/main"]