FROM golang:1.18-alpine
LABEL maintainer="Jan <jan.jiazhen@gmail.com>"

WORKDIR /
COPY . .
RUN apk --update --no-cache add g++
RUN go build -o block_api ./application/apiserver/main.go

EXPOSE 8081
ENTRYPOINT ["./block_api"]
