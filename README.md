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

Semver tags (`v0.1.0`). Pushing a tag triggers GoReleaser, which publishes binaries and GitHub Release notes.

`pkg.go.dev` picks up new tags automatically. Nothing to register.

## Changelog

See [CHANGELOG.md](./CHANGELOG.md), generated from conventional commits with [git-cliff](https://github.com/orhun/git-cliff):

```sh
git cliff --output CHANGELOG.md
```

## License

[MIT](./LICENSE)
