FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go get github.com/gorilla/mux github.com/google/uuid
RUN go mod tidy

COPY . ./

RUN go build -o app ./cmd/app/main.go

EXPOSE 8080

CMD ["./app"]
