package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"path/filepath"
)

func installSDK(srcDir, destDir string, copyBinaryFiles bool) {
	srcDir = filepath.Join(srcDir, "SDK")
	sdkDir := filepath.Join(destDir, "SDK")
	if copyAll {
		copySDK("", "", srcDir, sdkDir)
	} else {
		copySDK(goos, goarch, srcDir, sdkDir)
	}
	if copyBinaryFiles {
		copyBinaries(goos, goarch, srcDir, destDir)
	}
}

func copySDK(goos, goarch, srcDir, destDir string) {
	if strings.Contains(destDir, filepath.Join("vendor","github.com")) {
		fmt.Println("Copying SDK to vendor folder...")
	} else {
		fmt.Println("Copying SDK to project...")
	}
	var pathStub string
	switch goos {
	case "windows":
		if goarch == "amd64" {
			pathStub = filepath.Join("win", "x64")
		} else {
			pathStub = filepath.Join("win", "x86")
		}
	case "darwin":
		pathStub = "mac"
	case "linux":
		pathStub = "linux"
	}
	filepath.Walk(filepath.Join(srcDir, "bin", pathStub), func(path string, info os.FileInfo, err error) error {
		relPath, err := filepath.Rel(filepath.Join(srcDir, "bin"), path)
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
			return err
		}
		if info.IsDir() {
			os.MkdirAll(filepath.Join(destDir, "bin", relPath), 0777)
		} else {
			copy(path, filepath.Join(destDir, "bin", filepath.Dir(relPath)))
		}
		return nil
	})
	filepath.Walk(filepath.Join(srcDir, "include"), func(path string, info os.FileInfo, err error) error {
		relPath, err := filepath.Rel(filepath.Join(srcDir, "include"), path)
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
			return err
		}
		if info.IsDir() {
			os.MkdirAll(filepath.Join(destDir, "include", relPath), 0777)
		} else {
			copy(path, filepath.Join(destDir, "include", filepath.Dir(relPath)))
		}
		return nil
	})
	if goos == "windows" || goos == "" {
		filepath.Walk(filepath.Join(srcDir, "lib", pathStub), func(path string, info os.FileInfo, err error) error {
			relPath, err := filepath.Rel(filepath.Join(srcDir, "lib"), path)
			if err != nil {
				log.Fatalf("Failed to get relative path: %v", err)
				return err
			}
			if info.IsDir() {
				os.MkdirAll(filepath.Join(destDir, "lib", relPath), 0777)
			} else {
				copy(path, filepath.Join(destDir, "lib", filepath.Dir(relPath)))
			}
			return nil
		})
	}
	fmt.Println("Files copied!")
}

func copyBinaries(goos, goarch, srcDir, destDir string) {
	fmt.Println("Copying binaries...")
	var pathStub string
	switch goos {
	case "windows":
		if goarch == "amd64" {
			pathStub = filepath.Join("win", "x64")
		} else {
			pathStub = filepath.Join("win", "x86")
		}
	case "darwin":
		pathStub = "mac"
	case "linux":
		pathStub = "linux"
	}
	filepath.Walk(filepath.Join(srcDir, "bin", pathStub), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
			return err
		}
		if !info.IsDir() {
			copy(path, destDir)
		}
		return nil
	})
	fmt.Println("Files copied!")
}

func copy(srcFile, destDir string) {
	binData, err := ioutil.ReadFile(srcFile)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filepath.Base(srcFile), err)
	}
	err = ioutil.WriteFile(filepath.Join(destDir, filepath.Base(srcFile)), binData, 0777)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filepath.Base(srcFile), err)
	}
}
