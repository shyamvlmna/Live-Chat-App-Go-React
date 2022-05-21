FROM golang:1.18.2-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go mod download
ADD . /app/
RUN go build -o main ./
CMD ["/app/main"]