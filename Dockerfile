FROM golang:alpine as golang
RUN apk --update --no-cache add ca-certificates && update-ca-certificates
WORKDIR /go/src
COPY src/ .

RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' github.com/abrie/mapatlapi/cmd/cli

FROM scratch
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=golang /go/bin/cli /app
ENTRYPOINT ["/app"]



