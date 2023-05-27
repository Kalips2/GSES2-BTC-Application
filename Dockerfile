FROM golang:1.20-alpine AS build-stage

RUN apk --no-cache add ca-certificates

WORKDIR /btcApplication

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /btcApplication/btc-app .

RUN chmod +x /btcApplication/btc-app

FROM scratch

COPY --from=build-stage /btcApplication/btc-app /btcApplication/btc-app

COPY --from=build-stage /btcApplication/storage/ /btcApplication/storage

COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

ENTRYPOINT ["/btcApplication/btc-app"]