ARG GO_VERSION=1.23
ARG TARGET_OS=linux
ARG TARGET_ARCH=amd64

FROM --platform=$TARGET_OS/$TARGET_ARCH golang:$GO_VERSION AS builder

ARG TARGET_SERVICE


RUN if [ -z "$TARGET_SERVICE" ]; then echo "TARGET_SERVICE not exists" >&2; exit 1; fi

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN ./scripts/generate_sqlc.sh

RUN CGO_ENABLED=0 GOARCH=$TARGET_ARCH GOOS=$TARGET_OS go build -a -installsuffix cgo -o /bin/auth-service ./cmd/$TARGET_SERVICE

FROM alpine:latest

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

ENV CONFIG_PATH="/go/config/config.yaml"

RUN mkdir -p /go/config

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

COPY --from=builder /bin/auth-service /bin/

EXPOSE 9000

ENTRYPOINT [ "/bin/auth-service" ]
