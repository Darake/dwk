# syntax=docker/dockerfile:1

##
## Build server
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /go-app

##
## Build client
##
FROM node:12.22.1 AS node-build

WORKDIR /app

COPY ["package.json", "package-lock.json*", "./"]

RUN npm i

COPY public ./public
COPY src ./src

RUN npm run build

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-app /app
COPY --from=node-build /app/build /build

EXPOSE 8089

USER nonroot:nonroot

ENTRYPOINT ["/app"]