FROM golang:1.18-alpine
LABEL maintainer="Jan <jan.jiazhen@gmail.com>"

WORKDIR /
COPY . .
RUN apk --update --no-cache add g++
RUN go build -o indexer ./application/indexer/main.go

ENTRYPOINT ["./indexer"]
