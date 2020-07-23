FROM golang:1.14

WORKDIR /go/src/gratisdns-ddns-go
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["gratisdns-ddns-go"]
