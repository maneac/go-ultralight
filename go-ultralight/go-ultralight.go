package main

import (
	"flag"
	"log"
	"runtime"
	"strings"
)

type flagVals struct {
	dryRun   bool
	download bool
	copy     bool
	goos     string
	allOS    bool
	output   string
}

type utility struct {
	flags *flagVals
	// current working directory
	userDir   string
	vendorDir string
	isModule  bool
	// location of the go-ultralight package
	packageDir string
}

func osToUL(os string) string {
	switch os {
	case "linux":
		return "linux"
	case "darwin":
		return "mac"
	case "windows":
		return "win"
	}
	return ""
}

func main() {
	envOpts := flagVals{}
	flag.BoolVar(&envOpts.dryRun, "dry-run", false, "Performs a dry run.")
	flag.BoolVar(&envOpts.download, "download", true, "Download the SDK if it does not exist.")
	flag.BoolVar(&envOpts.copy, "copy", true, "Copy the binaries to your current directory.")
	flag.StringVar(&envOpts.goos, "os", runtime.GOOS, "Target OS to fetch the binaries for.\nSupported options: linux, windows, darwin")
	flag.BoolVar(&envOpts.allOS, "all", false, "Fetch binaries for all supported OSes. Overrides 'os' flag.")
	flag.StringVar(&envOpts.output, "output", "", "Copy binaries to a specific folder instead of the current directory.")
	flag.Parse()

	envOpts.goos = strings.TrimSpace(strings.ToLower(envOpts.goos))
	switch envOpts.goos {
	case "linux", "darwin", "windows":
	default:
		log.Fatalf("Unsupported OS specified: %v", envOpts.goos)
	}

	u := utility{
		flags: &envOpts,
	}

	// retrieves and configures the folders used by the utility
	u.getDirectories()

	// download and extract the SDK
	if u.flags.download {
		u.download()
	}

	// copy the binaries to the target project
	if u.flags.copy {
		u.install()
	}
}
