# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS base
WORKDIR /src
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod/ \
--mount=type=bind,source=go.sum,target=go.sum \
--mount=type=bind,source=go.mod,target=go.mod \
go mod download -x
RUN go get -u all
RUN go mod tidy

FROM base AS build-server
RUN --mount=type=cache,target=/go/pkg/mod/ \
--mount=type=bind,target=. \
go build -o /bin/server ./cmd/api

EXPOSE 8000

FROM scratch AS server
COPY --from=build-server /bin/server /bin/
ENTRYPOINT [ "/bin/server" ]
