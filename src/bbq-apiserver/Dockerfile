FROM golang:1.11-alpine as build_base

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/ssargent/bbq/bbq-apiserver

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN git --version
RUN go mod download

FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' .

FROM alpine
RUN apk add ca-certificates
WORKDIR /app
COPY --from=server_builder /go/bin/bbq-apiserver /app/bbq-apiserver
COPY --from=server_builder /go/src/github.com/ssargent/bbq/bbq-apiserver/migrations /app/migrations
EXPOSE 21337
ENTRYPOINT ["./bbq-apiserver"]