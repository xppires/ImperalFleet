FROM golang:1.23-alpine AS baseapp

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app


COPY go.mod ./
RUN go mod download
RUN go get -u github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux

COPY . .

RUN go mod tidy

FROM  baseapp
RUN go build  -mod=mod -o main ./cmd/api/main.go



EXPOSE 8080

CMD ["./main"]
