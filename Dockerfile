FROM golang:1.12 as builder

RUN mkdir -p /go/src/github.com/gleez/drone-kubernetes
WORKDIR /go/src/github.com/gleez/drone-kubernetes

# Force the go compiler to use modules
ENV GO111MODULE=on

COPY . .

# And compile the project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -extldflags -static" -a -tags netgo -installsuffix netgo

# And compile the project
# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'


FROM scratch
COPY --from=builder /go/src/github.com/gleez/drone-kubernetes/drone-kubernetes /drone-kubernetes

CMD ["/drone-kubernetes"]
