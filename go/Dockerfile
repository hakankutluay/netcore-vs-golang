FROM golang:1.22 AS builder

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o /api .

FROM busybox:1.36

COPY --from=builder /api /api

CMD ["/api"]
