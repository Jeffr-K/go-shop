#Dockerfile
FROM golang:1.19-alpine

RUN apk update && apk upgrade && apk add git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 4000

CMD ["go","run","./cmd/main.go"]