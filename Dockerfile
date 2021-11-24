FROM golang:1.16 as builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
EXPOSE 8080

CMD ["make", "run"]