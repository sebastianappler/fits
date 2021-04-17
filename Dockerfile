FROM golang:alpine AS builder

WORKDIR /
COPY . .
RUN mkdir from
RUN mkdir to
RUN go build -v -ldflags="-s -w"

FROM scratch
ENV FITS_ENVIRONMENT=docker

COPY --from=builder /fits /config.toml /from /to /

ENTRYPOINT ["/fits"]
