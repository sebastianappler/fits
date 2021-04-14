FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o fits .

WORKDIR /dist
RUN cp -r /build/. .

FROM scratch
COPY --from=builder /dist/fits /
COPY --from=builder /dist/config.toml /

ENTRYPOINT ["/fits"]
