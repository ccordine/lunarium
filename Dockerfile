# syntax=docker/dockerfile:1

FROM node:22-alpine AS web-build
WORKDIR /src/web
COPY web/package.json web/package-lock.json ./
RUN npm ci --ignore-scripts
COPY web/ ./
RUN npm run build

FROM golang:1.24-alpine AS go-build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
COPY internal/ ./internal/
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/lunarium .

FROM alpine:3.22
RUN addgroup -S -g 10001 lunarium && adduser -S -D -H -u 10001 -G lunarium lunarium
WORKDIR /app
COPY --from=go-build --chown=lunarium:lunarium /out/lunarium /app/lunarium
COPY --from=web-build --chown=lunarium:lunarium /src/web/dist /app/web/dist

ENV TZ=UTC
EXPOSE 8080
USER lunarium

HEALTHCHECK --interval=15s --timeout=3s --start-period=5s --retries=4 \
  CMD wget -qO- http://127.0.0.1:8080/api/v1/health >/dev/null || exit 1

ENTRYPOINT ["/app/lunarium"]
CMD ["-addr", ":8080", "-assets", "/app/web/dist"]

