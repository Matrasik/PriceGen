FROM golang:1.19-alpine as build
WORKDIR /build
COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o /main .

EXPOSE 8888

ENTRYPOINT ["/main"]
