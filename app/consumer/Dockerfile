# Pull latest golang image
FROM golang:1.15.6-alpine as build

# Install git
RUN apk add --no-cache git 

# Working directory path
WORKDIR /src

COPY go.mod /src
COPY go.sum /src

RUN go mod download

COPY consumer.go /src

RUN go build consumer.go

FROM alpine:latest  as runtime 

COPY --from=build /src/consumer /app/consumer

CMD ["/app/consumer"]
