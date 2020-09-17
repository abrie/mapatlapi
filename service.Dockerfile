FROM golang:alpine as golang
WORKDIR /go/src
COPY src/ .
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' github.com/abrie/mapatlapi/cmd/service

FROM scratch
COPY --from=golang /go/bin/service /app
ENTRYPOINT ["/app"]



