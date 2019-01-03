FROM alpine:3.8
ENV GOPATH=/root/go
RUN apk add go musl-dev
WORKDIR /root/go/src/frontend
COPY *.go ./
RUN go build
ENTRYPOINT ["/root/go/src/frontend/frontend"]
