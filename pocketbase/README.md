# Pocketbase

## Default Workflow

By default Pocketbase can be run simply with `pocketbase serve` (may need execute permissions).

## Extending Pocketbase

Pocketbase can be extended via a custom Go app, as documented in the [Pocketbase Docs](https://pocketbase.io/docs/use-as-framework/). The documentation is very minimal on creating a Go project though...

### Configure Go

Go can either be installed directly, or via [Go Version Manager](https://github.com/moovweb/gvm). If using GVM, the first installed version will likely need to be built from a binary (`-B` flag). Go installation (`$GOROOT`) must be available in path for developer tooling (VS Code, etc).

Once Go is installed, a module can be created with `go mod init` (already handled in project). Dependencies are installed with `go mod tidy`. Packages can be added with `go get <package_name>`.

### Development

A custom Pocketbase package can be run (in development) with `make serve-dev` or as a regular binary after building with `make build`. Specifying an output file name can be helpful (`go build -o pocketbase main.go`).

### Database

Pocketbase uses [`dbx`](https://github.com/pocketbase/dbx#readme) for database queries.
