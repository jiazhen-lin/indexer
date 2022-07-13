FROM golang:1.18-alpine
LABEL maintainer="Jan <jan.jiazhen@gmail.com>"

WORKDIR /
COPY . .
RUN go build -o migration ./script/migrate/main.go

ENTRYPOINT ["./migration"]
