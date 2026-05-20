# Agent Instructions

## Build Tools

Use [mise](https://mise.jdx.dev) for tool management and task running.

```bash
mise install        # install pinned tools
mise run build      # full CI build
mise tasks          # list all tasks
```

## Package Manager

Use **Go Modules**: `go mod tidy`, `go test ./...`

## Structure

- One API operation per file: `client_<collection>_<operation>.go`
- Client returns protobuf messages — see `proto/` for schemas
- Protobuf conversion functions live in `proto.go`

## CLI Architecture

The CLI is split into two layers to keep credential storage pluggable:

```
cli/
├── cli.go       # Store interface, Credentials, FileStore, Options
└── command.go   # NewCommand() — full command tree
cmd/rfms/
└── main.go      # Thin wrapper: wires FileStore to XDG paths
```

- `cli.Store` — interface with `Read(any)`, `Write(any)`, `Clear()` methods
- `cli.NewCommand(...Option)` — builds the Cobra command tree; receives stores via functional options (`WithCredentialStore`, `WithTokenStore`)
- `cmd/rfms/main.go` — only wires `FileStore` instances and calls `cli.NewCommand()`

This separation lets consumers embed the CLI in a larger tool or swap the storage backend (e.g. use an in-memory store in tests, or a keychain-backed store) without forking.

### Multi-Provider Authentication

The rFMS standard is implemented by multiple OEMs, each with their own auth:

- **Scania** — OAuth2 client credentials (`auth login scania --client-id --client-secret`). Token cached in the token store.
- **Volvo Trucks** — HTTP Basic Auth (`auth login volvo-trucks --username --password`). No token store needed.

The `Credentials` struct stores the active provider alongside the provider-specific fields.

### Embedding in a Parent CLI

The CLI can be embedded as a subcommand in a larger tool (e.g. a unified `way` CLI). Key design rules:

- **Never use `cmd.Root()`** — resolves to the parent CLI's root when embedded, breaking flag lookups. Use `cmd.Flags()` instead (works for both persistent and local flags).
- **`WithHTTPClient`** — the parent injects an `*http.Client` via `cli.WithHTTPClient()`. The SDK layers (auth, retry) stack on top of the injected client's transport.
- **`DebugTransport`** — exported in `debug.go` with a lazy `Enabled *bool` field. The parent owns the `--debug` flag and points `Enabled` at the flag variable. The transport checks the pointer at request time, solving the chicken-and-egg problem (transport constructed before flag parsing).

```go
var debug bool
cmd := cli.NewCommand(
    cli.WithCredentialStore(store),
    cli.WithTokenStore(tokenStore),
    cli.WithHTTPClient(&http.Client{
        Transport: &rfms.DebugTransport{
            Enabled: &debug,
            Next:    http.DefaultTransport,
        },
    }),
)
cmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
```

### Module Structure

Three separate Go modules prevent Cobra/CLI dependencies from leaking into the SDK library:

```
go.mod          # SDK client library (no cobra, no CLI deps)
cli/go.mod      # CLI commands (depends on root SDK + cobra)
cmd/rfms/go.mod # Standalone binary (depends on cli module)
```

Consumers who only need the Go client import the root module without pulling in CLI dependencies.

### Conventions

- Subcommands are organized by entity using `cobra.Group`
- Fully paginate results where applicable
- Flat command structure: `vehicles`, `vehicle-positions`, not `vehicles list`
