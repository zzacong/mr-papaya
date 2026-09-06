# Contributing

## Setup

- Go 1.24+
- Fork, branch, open a PR against `main`. CI runs `go build`, `go vet`, `go test`.

## Commits

Conventional commits (required for the auto-generated changelog):

- `feat: ...` new feature
- `fix: ...` bug fix
- `docs: ...`, `refactor: ...`, `test: ...`, `perf: ...`, `chore: ...`

## Release (maintainers)

1. Merge to `main`, ensure CI is green.
2. Regenerate the changelog: `git cliff --output CHANGELOG.md`, commit as `chore: release vX.Y.Z prep` if changed.
3. Tag and push: `git tag vX.Y.Z && git push origin main vX.Y.Z`.
4. GitHub Actions runs GoReleaser: cross-platform binaries, checksums, release notes.
5. Verify: `go install github.com/zzacong/mr-papaya/cmd/mr-papaya@vX.Y.Z` and check `pkg.go.dev/github.com/zzacong/mr-papaya` updates.
