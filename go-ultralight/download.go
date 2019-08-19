package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func downloadAll(srcDir string) {
	downloadSDK("windows", srcDir)
	downloadSDK("darwin", srcDir)
	downloadSDK("linux", srcDir)
}

func downloadSDK(goos, srcDir string) {
	var binStr string
	switch goos {
	case "windows":
		binStr = "win"
	case "darwin":
		binStr = "mac"
	case "linux":
		binStr = "linux"
	}
	if _, err := os.Stat(filepath.Join(srcDir, "SDK", "bin", binStr)); os.IsNotExist(err) {
		var url string
		switch goos {
		case "windows":
			url = "https://ultralig.ht/ultralight-sdk-1.0-win.zip"
		case "linux":
			url = "https://ultralig.ht/ultralight-sdk-1.0-linux.tar.xz"
		case "darwin":
			url = "https://ultralig.ht/ultralight-sdk-1.0-mac-v2.zip"
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to get SDK: %v", err)
		}

		fmt.Printf("Downloading %s Ultralight SDK...\n", goos)
		outDir := filepath.Join(srcDir, "_tmp")
		err = os.Mkdir(outDir, 0777)
		if err != nil {
			log.Fatalf("Failed to create temporary directory: %v", err)
		}
		out, err := os.Create(filepath.Join(outDir, filepath.Base(url)))
		if err != nil {
			log.Fatalf("Failed to create SDK file: %v", err)
		}
		defer os.RemoveAll(outDir)

		totalLen, err := strconv.Atoi(resp.Header.Get("Content-Length"))
		if err != nil {
			log.Printf("Failed to convert content-size %s: %v\n", resp.Header.Get("Content-Length"), err)
		}
		downloadCounter := &progress{totalSize: uint64(totalLen)}
		_, err = io.Copy(out, io.TeeReader(resp.Body, downloadCounter))
		if err != nil {
			log.Fatalf("Failed to write data to SDK file: %v", err)
		}
		fmt.Println()
		resp.Body.Close()
		out.Close()
		fmt.Println("SDK downloaded!")

		os.MkdirAll(filepath.Join(srcDir, "SDK"), 0777)

		if goos == "linux" {
			untar(filepath.Join(outDir, filepath.Base(url)), filepath.Join(srcDir, "SDK"), goos)
		} else {
			unzip(filepath.Join(outDir, filepath.Base(url)), filepath.Join(srcDir, "SDK"), goos)
		}
	}
}
