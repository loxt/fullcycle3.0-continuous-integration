FROM golang:1.18.3

WORKDIR /app

COPY . .

RUN go build -o math

CMD ["./math"]