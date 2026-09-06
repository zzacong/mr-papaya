# Contributing

## Setup

- Go 1.24+
- Fork, branch, open a PR against `main`. CI runs `go build`, `go vet`, `go test`.

## Commits

Conventional commits (required for the auto-generated changelog):

- `feat: ...` new feature
- `fix: ...` bug fix
- `docs: ...`, `refactor: ...`, `test: ...`, `perf: ...`, `chore: ...`

## Release (automated)

Merging to `main` is the only manual step. Release-please opens a rolling release PR (version bump + `CHANGELOG.md` from conventional commits). Merging that PR creates the tag and GitHub release; GoReleaser then publishes binaries.

Setup required once: create a fine-grained PAT with Contents read/write on this repo, save it as the `RELEASE_PLEASE_TOKEN` secret. Tags created with the default `GITHUB_TOKEN` cannot trigger the GoReleaser workflow, which is why the PAT exists.
