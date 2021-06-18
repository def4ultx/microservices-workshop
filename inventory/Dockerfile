FROM golang:1.14 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

ADD . /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/api .

FROM alpine

COPY --from=builder /build/api .
COPY --from=builder /build/config.yml config.yml
ENTRYPOINT ["./api"]
EXPOSE 8080