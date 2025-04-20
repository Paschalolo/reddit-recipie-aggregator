FROM golang:1.24.2
WORKDIR /go/src/github.com/worker
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build  -o app cmd/consumer/*.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/worker/app .
CMD ["./app"]
