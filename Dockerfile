FROM golang:alpine as golang
WORKDIR /go/src
COPY src/ .
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' github.com/abrie/mapatlapi/cmd/cli

FROM scratch
COPY --from=golang /go/bin/cli /app
ENTRYPOINT ["/app"]



