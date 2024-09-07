FROM golang:1.22.1-alpine3.18 AS builder
# FROM golang:1.19.5-alpine3.16 AS builder
RUN apk --no-cache add gcc g++ make git
RUN apk --no-cache add tzdata

# ARG CGO_ENABLED=0

WORKDIR /go/src/app

COPY . .

RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./cmd/main.go

# Start from a minimal Alpine image
FROM alpine:3.13

# # Install required dependencies
RUN apk --no-cache add ca-certificates

# # Set the working directory
WORKDIR /usr/bin

# RUN mkdir logging

# Copy the built Golang application from the builder stage
COPY --from=builder /go/src/app/bin /go/bin

COPY .env /usr/bin

ENV TZ=Asia/Bangkok 

EXPOSE 80
ENTRYPOINT /go/bin/web-app --port 80

