# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.21.7-alpine3.18 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
COPY academprojex-firebase-adminsdk.json ./academprojex-firebase-adminsdk.json  
RUN go build -v -o /usr/local/bin/app .

# Final stage
FROM alpine:3.14

COPY --from=build /usr/local/bin/app /usr/local/bin/app
COPY .env ./.env
COPY academprojex-firebase-adminsdk.json ./academprojex-firebase-adminsdk.json 

# Set the default environment variable for MONGOURI
ENV MONGOURI="mongodb://root:example@localhost:27017/"

EXPOSE 8080

CMD ["app"]
