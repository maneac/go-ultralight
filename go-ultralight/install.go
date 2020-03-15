package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var errInvalidSDK = errors.New("invalid SDK format")

func copySDK(goos, srcDir, destDir string) error {
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		fmt.Printf("\nNo SDK detected. Please download the Ultralight SDK from here:\n\n\thttps://github.com/ultralight-ux/Ultralight/releases\n\nthen extract it into the following folder:\n\n\t%v\n\nExample structure:\n\ngo-ultralight\n  |-SDK\n    |-bin\n    |-include\n\n", srcDir)
		return errInvalidSDK
	}
	if _, err := os.Stat(filepath.Join(srcDir, "bin")); os.IsNotExist(err) {
		fmt.Printf("\nIncorrect SDK file structure detected! Ensure the binary files can be found in:\n\n\t%v\n\nExample structure:\n\ngo-ultralight\n  |-SDK\n    |-bin\n    |-include\n\n", filepath.Join(srcDir, "bin"))
		return errInvalidSDK
	}
	if strings.Contains(destDir, filepath.Join("vendor", "github.com")) {
		fmt.Println("Copying SDK to vendor folder...")
	}
	filepath.Walk(filepath.Join(srcDir, "bin"), func(path string, info os.FileInfo, err error) error {
		relPath, err := filepath.Rel(filepath.Join(srcDir, "bin"), path)
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
			return err
		}
		if !info.IsDir() {
			switch goos {
			case "darwin":
				if filepath.Ext(path) == ".dylib" {
					copy(path, filepath.Join(destDir, "bin", filepath.Dir(relPath)))
				}
			case "linux":
				if filepath.Ext(path) == ".so" {
					copy(path, filepath.Join(destDir, "bin", filepath.Dir(relPath)))
				}
			case "windows":
				if filepath.Ext(path) == ".dll" {
					copy(path, filepath.Join(destDir, "bin", filepath.Dir(relPath)))
				}
			default:
				copy(path, filepath.Join(destDir, "bin", filepath.Dir(relPath)))
			}
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
		filepath.Walk(filepath.Join(srcDir, "lib"), func(path string, info os.FileInfo, err error) error {
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
	fmt.Println("SDK files copied!")
	return nil
}

func copyBinaries(goos, srcDir, destDir string) error {
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		fmt.Printf("\nNo SDK detected. Please download the Ultralight SDK from here:\n\n\thttps://github.com/ultralight-ux/Ultralight/releases\n\nthen extract it into the following folder:\n\n\t%v\n\nExample structure:\n\ngo-ultralight\n  |-SDK\n    |-bin\n    |-include\n\n", srcDir)
		return errInvalidSDK
	}
	if _, err := os.Stat(filepath.Join(srcDir, "bin")); os.IsNotExist(err) {
		fmt.Printf("\nIncorrect SDK file structure detected! Ensure the binary files can be found in:\n\n\t%v\n\nExample structure:\n\ngo-ultralight\n  |-SDK\n    |-bin\n    |-include\n\n", filepath.Join(srcDir, "bin"))
		return errInvalidSDK
	}
	fmt.Println("Copying binaries...")
	filepath.Walk(filepath.Join(srcDir, "bin"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
			return err
		}
		if !info.IsDir() {
			switch goos {
			case "darwin":
				if filepath.Ext(path) == ".dylib" {
					copy(path, destDir)
				}
			case "linux":
				if filepath.Ext(path) == ".so" {
					copy(path, destDir)
				}
			case "windows":
				if filepath.Ext(path) == ".dll" {
					copy(path, destDir)
				}
			default:
				copy(path, destDir)
			}
		}
		return nil
	})
	fmt.Println("Binary files copied!")
	return nil
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
