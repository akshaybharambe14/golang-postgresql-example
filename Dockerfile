# ------------------------------------------------------------------- build ---

# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:1.19-alpine as builder

LABEL maintainer="Akshay Bharambe <akshaybharambe14@gmail.com>"

WORKDIR /app

COPY . ./

RUN go mod download
RUN go build -o golang-postgresql-example

# ------------------------------------------------------------------- execute ---
FROM alpine:latest

# Setup target. More info at https://github.com/gliderlabs/docker-alpine/blob/master/docs/usage.md
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/golang-postgresql-example .

# Expose port 8537 to the outside world
EXPOSE 8537

# Command to run the executable
CMD ["./golang-postgresql-example"] 