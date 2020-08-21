FROM golang:1.14-alpine as server-builder

ENV GO111MODULE=on

RUN apk add --no-cache curl bash git protobuf-dev
RUN go get -u github.com/golang/protobuf/protoc-gen-go

COPY . /build/
WORKDIR /build/

RUN go generate grpcgen/gen.go

RUN go build -o grpc-hello-world

################################################################################

FROM golang:1.14-alpine as client-builder

ENV GO111MODULE=on

RUN apk add --no-cache curl bash git protobuf-dev
RUN go get -u github.com/golang/protobuf/protoc-gen-go

COPY . /build/
WORKDIR /build/

RUN go generate grpcgen/gen.go

RUN go build -o grpc-hello-world-client ./cmd/client

################################################################################

FROM alpine:3.12 as client

# Add certificates for SSL requests
RUN apk add --no-cache ca-certificates bash

WORKDIR /app
COPY --from=client-builder /build/grpc-hello-world-client ./

CMD [ "/app/grpc-hello-world-client" ]

################################################################################

FROM alpine:3.12 as server

# Add certificates for SSL requests
RUN apk add --no-cache ca-certificates bash

WORKDIR /app
COPY --from=server-builder /build/grpc-hello-world ./

CMD [ "/app/grpc-hello-world" ]
