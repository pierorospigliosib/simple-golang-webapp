FROM golang:1.23.6-alpine3.21
WORKDIR /app
COPY *.go .
RUN go mod init app
RUN go mod tidy
RUN go build
EXPOSE 8080
CMD ["./app"]
