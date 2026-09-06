# GoReleaser for releases

Releases are tag-driven: pushing `v*.*.*` runs GoReleaser via GitHub Actions to build cross-platform binaries and publish GitHub Release notes. Chosen over hand-rolled `go build` scripts because cross-compilation, archives, and checksums are boilerplate we don't want to maintain.

## Considered Options

- Hand-rolled goreleaser-free workflow with `go build` per platform.
