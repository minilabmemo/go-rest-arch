
#Multi-stage builds

## Build

FROM golang:1.18-alpine AS builder

ENV GO111MODULE=on

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base
WORKDIR /minilabmemo/go-rest-arch

# Copy go.mod and go.sum files to workspace  CACHED

COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy the source code
COPY . ./
#RUN ls 
WORKDIR /minilabmemo/go-rest-arch/cmd/app-core
# Build the Go app
RUN  make go-build


## Deploy
FROM alpine

COPY --from=builder /minilabmemo/go-rest-arch/cmd/app-core/app-core /app-core
COPY --from=builder /minilabmemo/go-rest-arch/cmd/app-core/configs/docker/service.toml /configs/docker/service.toml

ENTRYPOINT ["/app-core","--env=docker","--confdir=/configs"]

#Building 30.4s (18/18) FINISHED  
#Building 0.9s (18/18) FINISHED  CACHED