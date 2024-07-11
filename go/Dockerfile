FROM golang:1.22

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o /api .

CMD ["/api"]