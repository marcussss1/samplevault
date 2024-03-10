# syntax=docker/dockerfile:1
FROM golang:1.21.4-alpine AS base
WORKDIR /src
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod/ \
--mount=type=bind,source=go.sum,target=go.sum \
--mount=type=bind,source=go.mod,target=go.mod \
go mod download -x

FROM base AS build-server
RUN --mount=type=cache,target=/go/pkg/mod/ \
--mount=type=bind,target=. \
go build -tags go_tarantool_ssl_disable,go_tarantool_call_17 -o /bin/server ./cmd/api

EXPOSE 8000

FROM scratch AS server
COPY --from=build-server /bin/server /bin/
# TODO COPY /BIN/SH
ENTRYPOINT [ "/bin/server" ]
