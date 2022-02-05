### build stage ###
FROM golang:1.14-alpine AS builder
COPY . .
RUN go build -o server .


### run stage ###
FROM alpine:3.12
COPY --from=builder /go/server ./server
COPY web/index.html index.html
EXPOSE 8080
CMD ["./server"]
