#Dockerfile
FROM golang:1.19-alpine

WORKDIR /app

RUN apk update && \
    apk add git && \
    go install github.com/cespare/reflex@latest # Hot Reload module

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]
