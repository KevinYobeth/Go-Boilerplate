FROM golang:1.21-alpine AS golang

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server .

FROM gcr.io/distroless/static-debian12

COPY --from=golang /server .
COPY .env .

CMD ["/server"]