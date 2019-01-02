FROM golang:1
WORKDIR /go/src/frontend
COPY *.go ./
RUN go build
CMD ["/go/src/frontend/frontend"]
