FROM golang:1.21.1-alpine3.18

WORKDIR /go-cloud-native-rest-api
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/go-cloud-native-rest-api/bin/api"]
EXPOSE 8080