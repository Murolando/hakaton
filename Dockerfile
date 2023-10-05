FROM golang:1.19-alpine

RUN go version
WORKDIR /hakaton
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./

RUN go build -o hakaton ./cmd/main.go

CMD ["./hakaton"]