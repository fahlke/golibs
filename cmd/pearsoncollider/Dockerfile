# compile pearsoncollider binary
FROM registry.hub.docker.com/library/golang:1.13.5-buster as builder
LABEL stage=intermediate
COPY . /build
WORKDIR /build
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on
RUN go build \
      -ldflags="-w -s" \
      -o pearsoncollider \
      main.go

# build container image with the pearsoncollider binary
FROM scratch

LABEL maintainer  "Alexander Fahlke <alexander@fahlke.dev>"
LABEL description "An image running a small tool to show hash collisions for an arbitrary input string against a random generated string in my Pearson-16 hash implementation"

STOPSIGNAL SIGTERM
EXPOSE 3001

COPY --from=builder /build/pearsoncollider /pearsoncollider

ENV TEST_STRING="foo"

ENTRYPOINT ["/pearsoncollider"]