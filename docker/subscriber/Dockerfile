FROM golang:1.24.2-alpine AS golang

ARG SCRIPT_DIR

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server .$SCRIPT_DIR

FROM gcr.io/distroless/static-debian12

COPY --from=golang /server .
COPY .env .

CMD ["/server"]