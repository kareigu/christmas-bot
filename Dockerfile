FROM golang:alpine

WORKDIR /go/src/christmas_bot
COPY . .

RUN go get -d -v ./...
RUN apk add --update make
RUN make build

CMD ["bin/main"]