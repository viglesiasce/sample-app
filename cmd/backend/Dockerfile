FROM golang:1
RUN apt-get update && apt-get upgrade -y
WORKDIR /go/src/backend
COPY vendor/ ./vendor/
COPY *.go ./
RUN go build
CMD ["/go/src/backend/backend"]
