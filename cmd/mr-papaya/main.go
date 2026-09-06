package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const greeting = "hello, from mr papaya"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout))
}

func run(args []string, out io.Writer) int {
	fs := flag.NewFlagSet("mr-papaya", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	showVersion := fs.Bool("version", false, "print version and exit")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *showVersion {
		fmt.Fprintln(out, versionString())
		return 0
	}
	fmt.Fprintln(out, greeting)
	return 0
}

func versionString() string {
	if commit == "none" {
		return version
	}
	return fmt.Sprintf("%s (commit %s, built %s)", version, commit, date)
}
