FROM golang:1.20

## UPDATE THE OS
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y  && \
    apt-get install -y tzdata  

WORKDIR /go/src

## SET ENVIRONMENT
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on
ENV TZ America/Sao_Paulo

## PREPARE FOR protoc, gRPC & Evans
RUN apt-get install -y protobuf-compiler && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    wget https://github.com/ktr0731/evans/releases/download/0.9.1/evans_linux_amd64.tar.gz && \
    tar -xzvf evans_linux_amd64.tar.gz && \
    mv evans ../bin && rm -f evans_linux_amd64.tar.gz

## START A PROJECT
RUN go mod init server

## COPY NECESSARY FILES
# COPY go.mod go.sum ./
RUN go mod download && \
    go mod tidy

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]
