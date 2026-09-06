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
	for _, args := range [][]string{{"--version"}, {"-v"}} {
		var out bytes.Buffer
		if code := run(args, &out); code != 0 {
			t.Fatalf("args %v: exit code = %d, want 0", args, code)
		}
		if out.Len() == 0 {
			t.Fatalf("args %v: expected non-empty version output", args)
		}
	}
}

func TestRun_BadFlag(t *testing.T) {
	var out bytes.Buffer
	if code := run([]string{"--nope"}, &out); code == 0 {
		t.Fatal("expected non-zero exit for unknown flag")
	}
}
