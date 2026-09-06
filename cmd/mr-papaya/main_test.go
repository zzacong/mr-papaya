package main

import (
	"bytes"
	"testing"
)

func TestRun_PrintsGreeting(t *testing.T) {
	var out bytes.Buffer
	if code := run([]string{}, &out); code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if got, want := out.String(), greeting+"\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}

func TestRun_VersionFlag(t *testing.T) {
	var out bytes.Buffer
	if code := run([]string{"--version"}, &out); code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if out.Len() == 0 {
		t.Fatal("expected non-empty --version output")
	}
}

func TestRun_BadFlag(t *testing.T) {
	var out bytes.Buffer
	if code := run([]string{"--nope"}, &out); code == 0 {
		t.Fatal("expected non-zero exit for unknown flag")
	}
}
