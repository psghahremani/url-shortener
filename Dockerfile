FROM golang:1.15.8 AS build

RUN mkdir /app

COPY go.mod /app
COPY go.sum /app
WORKDIR /app

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM alpine:3.9

COPY --from=build /app/app app
RUN chmod +x /app

ENTRYPOINT ["/app"]
