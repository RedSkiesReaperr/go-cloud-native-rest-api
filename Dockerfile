FROM golang:1.21.1-alpine3.18

WORKDIR /go-cloud-native-rest-api
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/api \
    && go build -o ./bin/migrate ./cmd/migrate

CMD ["/go-cloud-native-rest-api/bin/api"]
EXPOSE 8080