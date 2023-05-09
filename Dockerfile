# FROM ngnix
# WORKDIR /app
# COPY go.mod .
# RUN go mod download
# COPY *.go .
# CMD ["./main"]
# EXPOSE 8080FROM golang:1.16-alpine AS builder
# WORKDIR /app
# COPY . .
# RUN go get github.com/gin-gonic/gin
# RUN go build -o main .
# EXPOSE 8080
# CMD ["./main"]


FROM httpd:2.4
WORKDIR /app
COPY ./machine.html /usr/local/apache2/htdocs/
COPY ./machine.css /usr/local/apache2/htdocs/
CMD ["/app/main"]
EXPOSE 80




