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

RUN mkdir from
RUN mkdir to

WORKDIR /dist
RUN cp -r /build/. .

FROM scratch
COPY --from=builder /dist/fits /
COPY --from=builder /dist/config.toml /
COPY --from=builder /dist/from /from
COPY --from=builder /dist/to /to

ENV FITS_ENVIRONMENT=docker

ENTRYPOINT ["/fits"]
