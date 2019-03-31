FROM golang:1.12 as builder

RUN mkdir -p /go/src/github.com/gleez/drone-kubernetes
WORKDIR /go/src/github.com/gleez/drone-kubernetes

COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -a 


FROM scratch
COPY --from=builder /go/src/github.com/gleez/drone-kubernetes/drone-kubernetes /drone-kubernetes

CMD ["/drone-kubernetes"]
