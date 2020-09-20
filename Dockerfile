FROM golang:1.15 AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN make compile

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app .
EXPOSE 8080
CMD ["./cmarket-linux_arm"]
