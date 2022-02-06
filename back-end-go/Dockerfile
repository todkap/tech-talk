### build stage ###
FROM golang:1.17.6-alpine AS builder

RUN apk --no-cache add ca-certificates

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app/back-end-gos

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
RUN go build -o server .


### run stage ###
FROM alpine:3.15
COPY --from=builder /go/server ./server
EXPOSE 8080
CMD ["./server"]
