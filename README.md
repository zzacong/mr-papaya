# mr-papaya

[![ci](https://github.com/zzacong/mr-papaya/actions/workflows/ci.yml/badge.svg)](https://github.com/zzacong/mr-papaya/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/zzacong/mr-papaya.svg)](https://pkg.go.dev/github.com/zzacong/mr-papaya)

Experimental Go CLI. Prints `hello, from mr papaya`.

## Install

Requires Go 1.24+.

```sh
go install github.com/zzacong/mr-papaya/cmd/mr-papaya@latest
```

Or download a binary from [Releases](https://github.com/zzacong/mr-papaya/releases).

## Usage

```sh
mr-papaya
mr-papaya --version
```

## Development

```sh
go run ./cmd/mr-papaya
go test ./...
go vet ./...
```

See [CONTRIBUTING.md](./CONTRIBUTING.md).

## Versioning

Automated with release-please: merging conventional commits to `main` opens a release PR; merging it tags `vX.Y.Z`, which triggers GoReleaser (binaries + notes).

`pkg.go.dev` picks up new tags automatically. Nothing to register.

## Changelog

See [CHANGELOG.md](./CHANGELOG.md), maintained by release-please from conventional commits. Do not edit by hand.

## License

[MIT](./LICENSE)
