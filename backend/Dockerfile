# STAGE 1 BUILD
FROM golang:1.19
WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /backend

EXPOSE 8000

CMD ["/backend"]
