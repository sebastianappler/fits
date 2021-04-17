FROM golang:alpine AS builder

WORKDIR /build
RUN mkdir from
RUN mkdir to

COPY . .
RUN go build -v -ldflags="-s -w"

FROM scratch
ENV FITS_ENVIRONMENT=docker

COPY --from=builder /build/fits /build/config.toml build/from build/to /

ENTRYPOINT ["/fits"]
