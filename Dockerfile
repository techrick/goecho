#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/goecho
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/goecho /goecho
ENTRYPOINT ./goecho
LABEL Name=goecho Version=0.0.1
EXPOSE 1323
