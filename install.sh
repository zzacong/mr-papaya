#!/bin/sh
# Install mr-papaya from GitHub Releases.
#
# Usage:
#   curl -sSL https://raw.githubusercontent.com/zzacong/mr-papaya/main/install.sh | sh
#   curl -sSL .../install.sh | sh -s -- v0.2.0   # specific version
#   INSTALL_DIR=$HOME/.local/bin sh install.sh   # custom directory
#
# Env:
#   INSTALL_DIR  destination directory (default: /usr/local/bin)
set -e

REPO="zzacong/mr-papaya"
BIN="mr-papaya"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

TAG="${1:-}"
if [ -z "$TAG" ]; then
	TAG="$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')"
fi
# GoReleaser asset versions strip the leading v: v0.2.0 -> 0.2.0
VER="$(printf '%s' "$TAG" | sed -E 's/^v//')"

OS="$(uname -s)"
case "$OS" in
	Linux) OS="linux" ;;
	Darwin) OS="darwin" ;;
	*)
		echo "unsupported OS: $(uname -s) (linux and darwin only)" >&2
		exit 1
		;;
esac

ARCH="$(uname -m)"
case "$ARCH" in
	x86_64) ARCH="amd64" ;;
	aarch64 | arm64) ARCH="arm64" ;;
	*)
		echo "unsupported architecture: $(uname -m) (amd64 and arm64 only)" >&2
		exit 1
		;;
esac

URL="https://github.com/$REPO/releases/download/$TAG/${BIN}_${VER}_${OS}_${ARCH}.tar.gz"

TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT INT TERM

echo "installing $BIN $TAG ($OS/$ARCH) to $INSTALL_DIR" >&2
curl -fsSL -o "$TMPDIR/$BIN.tar.gz" "$URL"
tar -xzf "$TMPDIR/$BIN.tar.gz" -C "$TMPDIR"
install -m 755 "$TMPDIR/$BIN" "$INSTALL_DIR/$BIN"
echo "installed: $("$INSTALL_DIR/$BIN" --version)" >&2
