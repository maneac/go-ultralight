package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func (u *utility) install() {
	if u.flags.allOS {
		for _, os := range []string{"linux", "darwin", "windows"} {
			u.copyFiles(os)
		}
		return
	}
	u.copyFiles(u.flags.goos)
}

func (u *utility) copyFiles(in string) {
	if u.flags.dryRun {
		if u.vendorDir != "" {
			log.Printf("Copy %s SDK to path: %s\n", in, u.vendorDir)
		}
		log.Printf("Copy %s binaries to path: %s\n", in, u.userDir)
		return
	}

	if _, err := os.Stat(filepath.Join(u.packageDir, "SDK")); os.IsNotExist(err) {
		log.Fatalf("\nNo SDK detected. Please download the Ultralight SDK using this utility, or manually from here:\n\n\thttps://github.com/ultralight-ux/Ultralight/releases\n\nthen extract it into the following folder:\n\n\t%v\n\nExample structure:\n\ngo-ultralight\n  |-SDK\n    |-bin\n    |-include\n\n", u.packageDir)
		return
	}
	if _, err := os.Stat(filepath.Join(u.packageDir, "SDK", "bin")); os.IsNotExist(err) {
		log.Fatalf("\nIncorrect SDK file structure detected! Ensure the binary files can be found in:\n\n\t%v\n\nExample structure:\n\ngo-ultralight\n  |-SDK\n    |-bin\n    |-include\n\n", filepath.Join(u.packageDir, "bin"))
	}

	if u.vendorDir != "" {
		u.copySDK(in)
	}

	u.copyBinaries(in)
}

func (u *utility) copySDK(in string) {
	log.Printf("Copying %s SDK to vendor folder...\n", in)
	filepath.Walk(filepath.Join(u.packageDir, "SDK", "bin"), func(path string, info os.FileInfo, err error) error {
		relPath, err := filepath.Rel(filepath.Join(u.packageDir, "bin"), path)
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
			return err
		}
		if !info.IsDir() {
			switch in {
			case "darwin":
				if filepath.Ext(path) == ".dylib" {
					copy(path, filepath.Join(u.vendorDir, "bin", filepath.Dir(relPath)))
				}
			case "linux":
				if filepath.Ext(path) == ".so" {
					copy(path, filepath.Join(u.vendorDir, "bin", filepath.Dir(relPath)))
				}
			case "windows":
				if filepath.Ext(path) == ".dll" {
					copy(path, filepath.Join(u.vendorDir, "bin", filepath.Dir(relPath)))
				}
			}
		}
		return nil
	})
	filepath.Walk(filepath.Join(u.packageDir, "SDK", "include"), func(path string, info os.FileInfo, err error) error {
		relPath, err := filepath.Rel(filepath.Join(u.packageDir, "include"), path)
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
			return err
		}
		if info.IsDir() {
			os.MkdirAll(filepath.Join(u.vendorDir, "include", relPath), 0777)
		} else {
			copy(path, filepath.Join(u.vendorDir, "include", filepath.Dir(relPath)))
		}
		return nil
	})
	if in == "windows" {
		filepath.Walk(filepath.Join(u.packageDir, "SDK", "lib"), func(path string, info os.FileInfo, err error) error {
			relPath, err := filepath.Rel(filepath.Join(u.packageDir, "lib"), path)
			if err != nil {
				log.Fatalf("Failed to get relative path: %v", err)
				return err
			}
			if info.IsDir() {
				os.MkdirAll(filepath.Join(u.vendorDir, "lib", relPath), 0777)
			} else {
				copy(path, filepath.Join(u.vendorDir, "lib", filepath.Dir(relPath)))
			}
			return nil
		})
	}
	log.Println("SDK files copied!")
}

func (u *utility) copyBinaries(in string) error {
	log.Printf("Copying %s binaries to project directory...\n", in)
	filepath.Walk(filepath.Join(u.packageDir, "SDK", "bin"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
			return err
		}
		if !info.IsDir() {
			switch in {
			case "darwin":
				if filepath.Ext(path) == ".dylib" {
					copy(path, u.userDir)
				}
			case "linux":
				if filepath.Ext(path) == ".so" {
					copy(path, u.userDir)
				}
			case "windows":
				if filepath.Ext(path) == ".dll" {
					copy(path, u.userDir)
				}
			}
		}
		return nil
	})
	log.Println("Binary files copied!")
	return nil
}

func copy(srcFile, destDir string) {
	binData, err := ioutil.ReadFile(srcFile)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filepath.Base(srcFile), err)
	}
	if _, err := os.Stat(destDir); err != nil {
		if err := os.MkdirAll(destDir, 0777); err != nil {
			log.Fatalf("Failed to make directory %s to copy files to: %v", destDir, err)
		}
	}
	err = ioutil.WriteFile(filepath.Join(destDir, filepath.Base(srcFile)), binData, 0777)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filepath.Base(srcFile), err)
	}
}
