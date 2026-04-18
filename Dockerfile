FROM golang:1.25-alpine AS builder

WORKDIR /src

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/app ./cmd/app

FROM alpine:3.22

RUN apk add --no-cache ca-certificates && adduser -D -u 10001 appuser

WORKDIR /app

COPY --from=builder /out/app /app/app
COPY migrations /app/migrations
COPY api /app/api

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/app"]
CMD ["api"]
