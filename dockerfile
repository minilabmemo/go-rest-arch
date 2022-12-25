FROM golang:1.18-alpine AS builder
RUN go env
WORKDIR /minilabmemo/go-rest-arch

# Copy go.mod and go.sum files to workspace
RUN pwd
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy the source code
COPY . ./

WORKDIR /minilabmemo/go-rest-arch/cmd/app-core
RUN go build -o app-core
#RUN ls  not working

FROM alpine
#WORKDIR /minilabmemo/go-rest-arch not working

COPY --from=builder /minilabmemo/go-rest-arch/cmd/app-core/app-core /app-core
COPY --from=builder /minilabmemo/go-rest-arch/cmd/app-core/configs/docker/service.toml /configs/docker/service.toml

ENTRYPOINT ["/app-core","--env=docker","--confdir=/configs"]