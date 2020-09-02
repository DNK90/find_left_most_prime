FROM golang:1.13
RUN mkdir -p $GOPATH/src/github.com/dnk90/find_left_most_prime
RUN mkdir -p $GOPATH/src/github.com/dnk90/find_left_most_prime/prime
WORKDIR $GOPATH/src/github.com/dnk90/find_left_most_prime
ADD go.mod main.go ./
ADD ./prime ./prime
RUN go get ./...
RUN go install
EXPOSE 8080
WORKDIR /go/bin
ENTRYPOINT ["./find_left_most_prime"]
