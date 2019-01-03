FROM launcher.gcr.io/google/debian9 AS builder
ENV GOVERSION=1.11.4 PATH=$PATH:/usr/local/go/bin GOPATH=/root/go
RUN apt-get update && apt-get install -y wget gcc && \
    wget -q https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz
WORKDIR /root/go/src/backend
COPY vendor/ ./vendor/
COPY *.go ./
RUN go build
FROM scratch
# Import the user and group files from the builder.
COPY --from=builder /root/go/src/backend/backend /backend
CMD ["/backend"]
