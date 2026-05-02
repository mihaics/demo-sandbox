# demo-sandbox

Minimal Go backend scaffold built with [Gin](https://github.com/gin-gonic/gin).
This is the foundation repo; subsequent feature work builds on top.

## Module path

```
github.com/mihaics/demo-sandbox
```

Import internal packages from this path (e.g.
`github.com/mihaics/demo-sandbox/internal/handlers`).

## Layout

```
cmd/api/            # service entrypoint (main package)
internal/server/    # router construction
internal/handlers/  # HTTP handlers
```

`cmd/` holds runnable binaries; `internal/` holds non-exported
packages so they cannot be imported by other modules.

## Requirements

- Go 1.25+ (the `go` directive in `go.mod` reflects the minimum;
  Gin pulls in this floor).

## Configuration

| Env var | Default  | Description                                                |
|---------|----------|------------------------------------------------------------|
| `PORT`  | `:8080`  | Listen address. Accepts `8080` or `:8080`; a bare port is prefixed with `:` automatically. |

## Run

```sh
make run
# or
go run ./cmd/api
```

The server listens on `$PORT` (default `:8080`) and exposes:

- `GET /healthz` — liveness probe, returns `{"status":"ok"}`.

Smoke test:

```sh
curl -fsS http://localhost:8080/healthz
```

## Build

```sh
make build
./bin/api
```

## Test

```sh
make test
# or
go test ./...
```

## Graceful shutdown

The server traps `SIGINT` and `SIGTERM` and shuts down via
`http.Server.Shutdown` with a 15-second timeout, allowing in-flight
requests to drain before exit.
