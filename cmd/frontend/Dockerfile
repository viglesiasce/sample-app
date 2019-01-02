FROM golang:1
RUN apt-get update && apt-get upgrade -y
WORKDIR /go/src/frontend
COPY *.go ./
RUN go build
CMD ["/go/src/frontend/frontend"]
