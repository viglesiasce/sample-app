FROM launcher.gcr.io/google/debian9
ENV GOVERSION=1.11.4 PATH=$PATH:/usr/local/go/bin
RUN apt-get update && apt-get install -y wget gcc && \
    wget -q https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOVERSION}.linux-amd64.tar.gz
WORKDIR /root/go/src/backend
COPY vendor/ ./vendor/
COPY *.go ./
RUN go build
CMD ["/root/go/src/backend/backend"]
