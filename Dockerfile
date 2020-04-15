FROM golang:1.13.6-alpine as builder

RUN apk add --update --no-cache git build-base musl-dev linux-headers

COPY . /build
WORKDIR /build

RUN go mod download
RUN go build -o bin/replacer .

FROM alpine:latest
COPY --from=builder /build/bin /bin/
COPY --from=builder /build/docs /docs/
COPY --from=builder /build/.env /bin/.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 3000

ENTRYPOINT ["/bin/replacer", "api"]
