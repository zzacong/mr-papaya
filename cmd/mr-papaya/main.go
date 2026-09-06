package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"
)

const greeting = "hello, from mr papaya"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, out, errOut io.Writer) int {
	fs := flag.NewFlagSet("mr-papaya", flag.ContinueOnError)
	fs.SetOutput(errOut)
	showVersion := fs.Bool("version", false, "print version and exit")
	showShortVersion := fs.Bool("v", false, "print version and exit (shorthand)")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *showVersion || *showShortVersion {
		fmt.Fprintln(out, versionString())
		return 0
	}
	fmt.Fprintln(out, greeting)
	return 0
}

func versionString() string {
	v := version
	// go install pkg@version builds from source without our ldflags,
	// so fall back to the module version embedded by the go tool.
	if v == "dev" {
		if bi, ok := debug.ReadBuildInfo(); ok && bi.Main.Version != "" {
			v = bi.Main.Version
		}
	}
	if commit == "none" {
		return v
	}
	return fmt.Sprintf("%s (commit %s, built %s)", v, commit, date)
}
