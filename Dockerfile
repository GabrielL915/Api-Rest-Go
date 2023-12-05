FROM golang:1.14.6-alpine3.12 AS builder
COPY go.mod go.sum /go/src/github.com/GabrielL915/Api-Rest-Go/
WORKDIR /go/src/github.com/GabrielL915/Api-Rest-Go
RUN go mod download
COPY . /go/src/github.com/GabrielL915/Api-Rest-Go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Api-Rest-Go github.com/GabrielL915/Api-Rest-Go

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/GabrielL915/Api-Rest-Go/build/Api-Rest-Go /usr/bin/Api-Rest-Go
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/Api-Rest-Go"]