FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN export PATH=$PATH:/go/bin

RUN migrate -database "mysql://root:password@tcp(mysql:3306)/go-marketplace" -path db/migrations/ up

EXPOSE 8080

CMD ["./main"]