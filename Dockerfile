FROM golang:1.20-buster

WORKDIR /app

COPY . .

RUN apt-get -y install wget git procps
RUN go mod tidy
RUN go mod download
RUN go build -o ./cmd/apiserver/apiserver .

CMD ["./cmd/apiserver/apiserver"]


