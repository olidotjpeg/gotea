FROM golang:1.22-bookworm AS dev

EXPOSE 8000

WORKDIR /build

# Install git and air so we can hot reload golang in testing mode
RUN apt-get update && \
    apt-get install -y git && \
    go install github.com/cosmtrek/air@latest

# Copy code and source files
COPY gotea.sqlite .
COPY linux.air.toml /air.toml
COPY internal ./internal
COPY cmd ./cmd

# Fetch go dependencies
COPY go.mod .
COPY go.sum .

RUN cat go.mod

RUN go mod tidy
RUN go mod download

ENTRYPOINT ["air", "-c", "/air.toml"]