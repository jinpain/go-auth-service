FROM golang:1.25.2-alpine3.22 AS build

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app ./cmd/

FROM alpine:3.22 AS final

COPY --from=build /app /bin/app

COPY --from=build /build/config/config.yaml /etc/app/config.yaml

COPY --from=build /build/sql/ /sql/

EXPOSE 8080

ENV CONFIG_PATH=/etc/app/config.yaml

ENTRYPOINT ["/bin/app"]