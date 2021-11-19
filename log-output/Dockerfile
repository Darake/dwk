# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /log-output

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /log-output /log-output

EXPOSE 5010

USER nonroot:nonroot

ENTRYPOINT ["/log-output"]