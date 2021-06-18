FROM golang:1.14 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

ADD . /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /build/consumer .

FROM alpine

COPY --from=builder /build/consumer .
ENTRYPOINT ["./consumer"]