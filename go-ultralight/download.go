package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/gen2brain/go-unarr"
)

type progress struct {
	totalSize uint64
	numBytes  uint64
}

func (p *progress) Write(data []byte) (int, error) {
	p.numBytes += uint64(len(data))
	p.Print()
	return len(data), nil
}

func (p *progress) Print() {
	fmt.Printf("\r\rProgress: %s/%s %s", humanize.Bytes(p.numBytes), humanize.Bytes(p.totalSize), strings.Repeat(" ", 10))
}

func (u *utility) download() {
	if u.flags.allOS {
		for _, os := range []string{"linux", "darwin", "windows"} {
			u.downloadSDK(os)
		}
		return
	}
	u.downloadSDK(u.flags.goos)
}

func (u *utility) downloadSDK(in string) {
	outDir := filepath.Join(u.packageDir, "SDK")
	if u.doesSDKExist(in) {
		return
	}
	if u.flags.dryRun {
		log.Printf("Download %s SDK to path: %s\n", in, outDir)
		return
	}
	resp, err := http.Get(fmt.Sprintf("https://github.com/ultralight-ux/Ultralight/releases/download/v1.1.0/ultralight-sdk-1.1.0-%s-x64.7z", osToUL(in)))
	if err != nil {
		log.Fatalf("Failed to get archive from URL: %v", err)
	}
	defer resp.Body.Close()
	totalLen, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		log.Printf("Failed to convert content-size %s: %v\n", resp.Header.Get("Content-Length"), err)
	}
	log.Printf("Downloading %s SDK...\n", in)
	downloadCounter := &progress{totalSize: uint64(totalLen)}
	a, err := unarr.NewArchiveFromReader(io.TeeReader(resp.Body, downloadCounter))
	if err != nil {
		log.Fatalf("Failed to open archive: %v", err)
	}
	defer a.Close()
	log.Printf("\nExtracting %s SDK to %s\n", in, outDir)
	_, err = a.Extract(outDir)
	if err != nil {
		log.Fatalf("Failed to extract archive: %v", err)
	}
	log.Println("Extraction complete!")
}

func (u *utility) doesSDKExist(in string) bool {
	if _, err := os.Stat(filepath.Join(u.packageDir, "SDK", "bin")); err != nil {
		return false
	}
	exist := false
	filepath.Walk(filepath.Join(u.packageDir, "SDK", "bin"), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		switch in {
		case "windows":
			if filepath.Ext(path) == ".dll" {
				exist = true
			}
		case "linux":
			if filepath.Ext(path) == ".so" {
				exist = true
			}
		case "darwin":
			if filepath.Ext(path) == ".dylib" {
				exist = true
			}
		}
		return nil
	})
	return exist
}
