# build stage
FROM golang:1.9.6-alpine3.7 AS build-env

RUN \
  apk update && \
  apk add git make

ADD Makefile /go/src/github.com/kublr/workshop-microservice-build-pipeline-colorer/Makefile
WORKDIR /go/src/github.com/kublr/workshop-microservice-build-pipeline-colorer

RUN make tools-update

ADD . /go/src/github.com/kublr/workshop-microservice-build-pipeline-colorer

RUN make deps-update

RUN make build

# final stage
FROM alpine:3.7
COPY --from=build-env /go/src/github.com/kublr/workshop-microservice-build-pipeline-colorer/target/server /opt/colorer/server
ENTRYPOINT ["/opt/colorer/server"]
EXPOSE 10000
