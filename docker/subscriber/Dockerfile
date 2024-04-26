FROM golang:alpine

ARG WORK_DIR

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary /app$WORK_DIR

RUN pwd

ENTRYPOINT ["/app/binary"]