FROM golang:1.25.2 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./cmd/main.go

# --- --- --- --- --- --- #
FROM gcr.io/distroless/base-debian12

COPY --from=build /app/app /app/app

COPY --from=build /app/config /app/config

COPY --from=build /app/sql /app/sql

EXPOSE 8080

ENTRYPOINT ["/app/app"]
