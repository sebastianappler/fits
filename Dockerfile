FROM golang:alpine AS builder

ENV CGO_ENABLED=0

WORKDIR /build
RUN mkdir config
RUN mkdir from
RUN mkdir to
RUN mkdir .ssh

COPY . .
RUN rm config/config.toml
RUN go build -v -ldflags="-s -w" ./cmd/fits

FROM scratch
ENV FITS_ENVIRONMENT=docker

COPY --from=builder /build/fits build/config build/from build/.ssh build/to /

ENTRYPOINT ["/fits"]
