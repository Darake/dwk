# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /

COPY go.mod ./
COPY *.go ./

RUN go build -o /app

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app /app

USER nonroot:nonroot

ENTRYPOINT ["/app"]