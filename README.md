<img src="doc/kontena-gopher.jpg?raw=true" title="Kontena Gopher" alt="Kontena Gopher" width="50%"> *(Gopher Logo CC BY 3.0: Renee French)*

# `github.com/kontena/kontena-client-go`

Golang Kontena API client library.

***DISCLAIMER***: This client library is a work in progress, and is intended for use with the Kontena 1.4 development version.

## Build

Requires Go 1.7.

```
  $ go get github.com/kontena/kontena-client-go
```

## Usage / Examples

### `./cli/commands`

Referr to the `cli/commands` package for examples showing how to implement the familar `kontena` CLI commands using the `client` and `api` package interface.

### `./cmd/kontena-cli`

An example implementation of the Kontena CLI using the `cli/commands` packages.

***This is not and is not intended to be a replacement for the official Kontena CLI. It is only intended for testing purposes, as an example application for using the `client` package!***

```
kontena-cli --help
NAME:
   kontena-cli - Kontena CLI

USAGE:
   kontena-cli [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     whoami   Show kontena master user
     grid     Kontena Grids
     node     Kontena Nodes
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug         [$KONTENA_DEBUG]
   --verbose      
   --quiet        
   --url value     [$KONTENA_URL]
   --token value   [$KONTENA_TOKEN]
   --help, -h     show help
   --version, -v  print the version
```

## Imports

### `github.com/kontena/kontena-client-go/api`

Kontena API definitions.

### `github.com/kontena/kontena-client-go/client`

Kontena Client struct, API interfaces.

Includes support for OAuth2 tokens and code exchanges.

## API Support

This client library is a work in progress, and support is limited to the following Kontena APIs:

* `/oauth2`
* `/v1/user`
* `/v1/grids`
* `/v1/nodes`

This library is currently compatible with the Kontena 1.4 development version.
