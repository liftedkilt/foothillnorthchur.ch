FROM golang:1.11.0-stretch as builder
WORKDIR /go-modules
COPY . ./
# Building using -mod=vendor, which will utilize the v
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o app -ldflags '-extldflags "-static"' ./apiserver.go

FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
# -0 means no compression.  Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -r -0 /zoneinfo.zip .

FROM scratch
# the test program:
COPY --from=builder /go-modules/app /app
# the timezone data:
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine /zoneinfo.zip /
# the tls certificates:
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app"]

EXPOSE 8888