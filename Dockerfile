FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

ENV GIN_MODE=release

RUN go build -o ./golibraryapi

EXPOSE 8080

CMD ["./golibraryapi"]
