package main

import (
	"log"
	"os"
	"path/filepath"
)

// getDirectories gets the required directories
func (u *utility) getDirectories() {
	var err error
	if u.flags.output != "" {
		u.userDir = u.flags.output
		if _, err := os.Stat(u.userDir); err != nil {
			log.Fatalf("Invalid output directory specified: %s", u.userDir)
		}
	} else {
		u.userDir, err = os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get current directory: %v", err)
		}
	}
	// check if directory is a Go Module project
	moduleFiles, err := filepath.Glob(filepath.Join(u.userDir, "go.mod"))
	if err != nil {
		log.Fatalf("Failed to check for module file: %v", err)
	}
	u.isModule = len(moduleFiles) > 0
	vendorFolder, err := filepath.Glob(filepath.Join(u.userDir, "vendor/"))
	if err != nil {
		log.Fatalf("Failed to check for vendor folder: %v", err)
	}
	if len(vendorFolder) > 0 {
		u.vendorDir = filepath.Join(u.userDir, "vendor", "github.com", "maneac", "go-ultralight")
	}
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalf("GOPATH not set!")
	}
	srcDir, err := filepath.Abs(gopath)
	if err != nil {
		log.Fatalf("Failed to get GOPATH directory path: %v", err)
	}
	if u.isModule {
		srcDir = filepath.Join(srcDir, "pkg", "mod", "github.com", "maneac")
		ulDir, err := filepath.Glob(filepath.Join(srcDir, "go-ultralight*"))
		if err != nil {
			log.Fatalf("Failed to search module package directory: %v", err)
		}
		if len(ulDir) == 0 {
			log.Fatalf("Failed to find go-ultralight in package directory!")
		}
		u.packageDir = ulDir[0]
		return
	}
	u.packageDir = filepath.Join(srcDir, "src", "github.com", "maneac", "go-ultralight")

	if _, err := os.Stat(u.packageDir); os.IsNotExist(err) {
		log.Fatalf("Package go-ultralight not found in %s!", u.packageDir)
	}
}
