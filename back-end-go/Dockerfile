### build stage ###
FROM golang:1.17.6-alpine AS builder

RUN mkdir /build
ADD *.go /build/
ADD *.mod /build/
ADD *.sum /build/

WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o build


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build .
EXPOSE 9080
CMD ["./build"]
