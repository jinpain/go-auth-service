FROM golang:1.25.2 AS build

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.26.0

# --- --- --- --- --- --- #
FROM gcr.io/distroless/base-debian12

COPY --from=build /go/bin/goose /

ENTRYPOINT [ "./goose" ]