FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY influx/*.go ./influx/
COPY ekz/*.go ./ekz/
RUN CGO_ENABLED=0 GOOS=linux go build -o /ekz-influx

CMD ["/ekz-influx"]