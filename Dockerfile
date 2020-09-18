FROM golang:alpine as golang
RUN apk --update --no-cache add ca-certificates && update-ca-certificates
WORKDIR /go/src
COPY src/ .

RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' github.com/abrie/mapatlapi/cmd/cli


# See https://stackoverflow.com/a/55757473/12429735
ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

FROM scratch
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=golang /go/bin/cli /app
COPY --from=golang /etc/passwd /etc/passwd
COPY --from=golang /etc/group /etc/group

# Use unprivileged, non-root user.
USER appuser:appuser

ENTRYPOINT ["/app"]
