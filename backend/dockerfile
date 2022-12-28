FROM golang:alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy
RUN go mod download

COPY . ./

RUN go build -o api ./...

EXPOSE 8080

CMD [ "./api", "8080" ]