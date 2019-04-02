FROM golang:1.12 as builder

RUN mkdir -p /go/src/drone-kubernetes
WORKDIR /go/src/drone-kubernetes

COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -a 


FROM scratch
COPY --from=builder /go/src/drone-kubernetes/drone-kubernetes /drone-kubernetes

CMD ["/drone-kubernetes"]
