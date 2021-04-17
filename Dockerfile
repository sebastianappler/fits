FROM golang:alpine AS builder

WORKDIR /fits
RUN mkdir from
RUN mkdir to

COPY . .
RUN go build -v -ldflags="-s -w"

FROM scratch
ENV FITS_ENVIRONMENT=docker

COPY --from=builder /fits/fits /fits/config.toml fits/from fits/to /

ENTRYPOINT ["/fits"]
