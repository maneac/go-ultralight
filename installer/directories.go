package main

import (
	"log"
	"os"
	"path/filepath"
)

func isModule(dir string) bool {
	moduleFiles, err := filepath.Glob(filepath.Join(dir, "go.mod"))
	if err != nil {
		log.Fatalf("Failed to check for module file: %v", err)
	}
	return len(moduleFiles) > 0
}

func isVendor(dir string) bool {
	vendorFoler, err := filepath.Glob(filepath.Join(dir, "vendor/"))
	if err != nil {
		log.Fatalf("Failed to check for vendor folder: %v", err)
	}
	return len(vendorFoler) > 0
}

func getSrcDir(isModule bool) string {
	srcDir, err := filepath.Abs(os.Getenv("GOPATH"))
	if err != nil {
		log.Fatalf("Failed to get GOPATH: %v", err)
	}
	if isModule {
		srcDir = filepath.Join(srcDir, "pkg", "mod", "github.com", "maneac")
		ulDir, err := filepath.Glob(filepath.Join(srcDir, "go-ultralight*"))
		if err != nil {
			log.Fatalf("Failed to search module package directory: %v", err)
		}
		if len(ulDir) > 0 {
			srcDir = ulDir[0]
		} else {
			log.Fatalf("Failed to find go-ultralight in package directory!")
		}
	} else {
		srcDir = filepath.Join(srcDir, "src", "github.com", "maneac", "go-ultralight")
	}

	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		log.Fatalf("go-ultralight not found!")
	}
	return srcDir
}

func getDestDir(outPath string) string {
	dest, err := filepath.Abs(outPath)
	if err != nil {
		log.Fatalf("Failed to get output directory: %v", err)
	}
	return dest
}

func getCurrentDir() string {
	dir, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	return dir
}
