FROM golang:alpine3.12
WORKDIR /go/src/github.com/abhide/simple-httpapp/
COPY main.go .
RUN go build -o simple-httpapp ./main.go

FROM alpine:3.12
WORKDIR /root/
COPY --from=0 /go/src/github.com/abhide/simple-httpapp/simple-httpapp .
CMD ["./simple-httpapp"]
