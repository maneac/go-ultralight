package main

import (
	"flag"
	"path/filepath"
	"runtime"
)

var goos string
var copyAll bool

func main() {
	flag.StringVar(&goos, "os", runtime.GOOS, "Target OS to fetch the binaries for")
	flag.BoolVar(&copyAll, "copy-all", false, "Copy the binaries for every OS")
	flag.Parse()

	curDir := getCurrentDir()

	srcDir := filepath.Join(getSrcDir(isModule(curDir)), "SDK")

	if copyAll {
		if isVendor(curDir) {
			if err := copySDK("", srcDir, filepath.Join(curDir, "vendor", "github.com", "maneac", "go-ultralight", "SDK")); err != nil {
				return
			}
		}
		copyBinaries("", srcDir, curDir)
		return
	}

	if isVendor(curDir) {
		if err := copySDK(goos, srcDir, filepath.Join(curDir, "vendor", "github.com", "maneac", "go-ultralight", "SDK")); err != nil {
			return
		}
	}
	copyBinaries(goos, srcDir, curDir)
}
