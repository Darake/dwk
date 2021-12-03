# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY *.go ./
ADD build ./build

RUN go build -o /go-app

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-app /app
COPY --from=build /app/build /build

EXPOSE 8089

USER nonroot:nonroot

ENTRYPOINT ["/app"]