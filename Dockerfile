FROM golang:1.20.4-alpine as build-stage

WORKDIR /app/btc-app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /btc-app .

FROM scratch

COPY --from=build-stage /btc-app /btc-app

EXPOSE 8080

ENTRYPOINT ["/btc-app"]