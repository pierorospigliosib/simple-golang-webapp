FROM golang:1.23.6-alpine3.21
COPY * .
RUN go build -o app .
EXPOSE 8080
CMD ["./app"]
