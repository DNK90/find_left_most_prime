FROM golang:1.13
RUN apt-get update && apt-get install -y unzip build-essential
# install protobuf
ENV PROTOC_ZIP=protoc-3.11.4-linux-x86_64.zip
WORKDIR /tmp
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/$PROTOC_ZIP
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN unzip -o $PROTOC_ZIP -d /usr/local include/*
RUN rm -f $PROTOC_ZIP
RUN ln -s /usr/local/bin/protoc /bin

# install protoc-gen-go
RUN go get github.com/gogo/protobuf/proto github.com/gogo/protobuf/protoc-gen-gogo

RUN mkdir -p $GOPATH/src/github.com/dnk90/find_left_most_prime
RUN mkdir -p $GOPATH/src/github.com/dnk90/find_left_most_prime/prime
WORKDIR $GOPATH/src/github.com/dnk90/find_left_most_prime
ADD go.mod main.go ./
ADD ./prime ./prime
ADD ./proto ./proto
# generate prime.proto
RUN protoc -I./proto --gogo_out=./proto ./proto/prime.proto
RUN go get ./...
RUN go install
EXPOSE 8080
WORKDIR /go/bin
ENTRYPOINT ["./find_left_most_prime"]
