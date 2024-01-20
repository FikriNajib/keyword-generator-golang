# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine3.16

RUN apk update && apk add --no-cache git

WORKDIR /app

# Download Go modules
COPY . .
RUN go mod tidy

RUN go build -o /keyword-generator ./src/cmd/rest


EXPOSE ${PORT}

# Run
CMD [ "/keyword-generator" ]

RUN go mod tidy
RUN go mod vendor


