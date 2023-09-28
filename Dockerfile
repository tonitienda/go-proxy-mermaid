ARG GO_ARCH=amd64

# Download Dependencies
FROM golang:alpine3.18 as dependencies

WORKDIR /usr/app/src

COPY proxy .

RUN go mod download


# Run as DEV
FROM dependencies as dev

WORKDIR /usr/app/src

RUN adduser -D appuser
USER appuser

CMD go run cmd/main.go

# Build Binary
FROM dependencies as build
WORKDIR /usr/app/src

RUN echo "Building for $GO_ARCH"

RUN GOOS=linux GOARCH=$GO_ARCH CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o /usr/app/src/proxy /usr/app/src/cmd/main.go && \
    mkdir -p /usr/app && \
    cp /usr/app/src/proxy /usr/app/proxy


# Runtime
FROM alpine:3.18 as prod

WORKDIR /usr/app/src
RUN adduser -D appuser
USER appuser

COPY --from=build /usr/app/src/proxy /usr/app/proxy

HEALTHCHECK --interval=30s --timeout=60s --retries=3 \
    CMD wget --no-verbose --tries=5 --spider http://localhost:$PORT/healthz || exit 1

CMD ENVIRONMENT=production GIN_MODE=release /usr/app/proxy
