FROM alpine:3.8
ENV GOPATH=/root/go
RUN apk add go musl-dev
WORKDIR /root/go/src/backend
COPY vendor/ ./vendor/
COPY *.go ./
RUN go build
CMD ["/root/go/src/backend/backend"]
