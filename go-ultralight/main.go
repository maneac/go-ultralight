package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var goos, goarch string
var offline, getAll, copyAll, download bool

func main() {
	flag.BoolVar(&download, "download", false, "Download SDKs only")
	flag.BoolVar(&offline, "offline", false, "Stops the tool from automatically downloading the Ultralight SDK")
	flag.BoolVar(&getAll, "all", false, "Download the Ultralight SDK for every OS")
	flag.StringVar(&goos, "os", runtime.GOOS, "Target OS to fetch the binaries for")
	flag.StringVar(&goarch, "arch", runtime.GOARCH, "Target OS to fetch the binaries for")
	flag.BoolVar(&copyAll, "targetAll", false, "Fetch the binaries for every OS")
	flag.Parse()

	curDir := getCurrentDir()

	if copyAll {
		binDir := filepath.Join(curDir, "bin")
		if _, err := os.Stat(binDir); os.IsNotExist(err) {
			err = os.MkdirAll(binDir, 0777)
			if err != nil {
				log.Fatalf("Failed to create binary directory: %v", err)
			}
		}
	}

	srcDir := getSrcDir(isModule(curDir))

	if !offline {
		if getAll {
			downloadAll(srcDir)
		} else {
			downloadSDK(goos, srcDir)
		}
	}

	if isVendor(curDir) {
		if copyAll {
			copySDK("", "", srcDir, filepath.Join(curDir, "vendor", "github.com", "maneac", "go-ultralight", "SDK"))
		} else {
			copySDK(goos, goarch, srcDir, filepath.Join(curDir, "vendor", "github.com", "maneac", "go-ultralight", "SDK"))
		}
	}

	if !download {
		if copyAll {
			copyBinaries("", "", filepath.Join(srcDir, "SDK"), curDir)
		} else {
			copyBinaries(goos, goarch, filepath.Join(srcDir, "SDK"), curDir)
		}
	}
}
