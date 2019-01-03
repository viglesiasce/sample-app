FROM alpine:3.8
ENV GOPATH=/root/go
RUN apk add go
WORKDIR /root/go/src/backend
COPY vendor/ ./vendor/
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CMD ["/root/go/src/backend/backend"]
