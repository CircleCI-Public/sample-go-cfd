FROM golang:1.26.10 AS build
WORKDIR /go/src
COPY go.mod go.sum ./
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go mod download

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
COPY --from=build /go/src/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
