FROM golang:1.24.3-alpine3.22 AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build ./cmd/shopapi

FROM alpine:3.22.0

WORKDIR /app
COPY --from=builder /build .

ENV RUNNING_IN_CONTAINER=true

RUN chmod +x ./shopapi

ENTRYPOINT [ "./shopapi" ]