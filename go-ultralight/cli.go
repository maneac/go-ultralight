package main

import (
	"archive/tar"
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/xi2/xz"
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

func getInstallationType() (bool, bool) {
	fmt.Println("Checking project type...")
	moduleFiles, err := filepath.Glob("go.mod")
	if err != nil {
		log.Fatalf("Failed to check for module file: %v", err)
	}
	isModule := len(moduleFiles) > 0

	vendorFolder, err := filepath.Glob("vendor/")
	if err != nil {
		log.Fatalf("Failed to check for vendor folder: %v", err)
	}
	isVendor := len(vendorFolder) > 0

	return isModule, isVendor
}

func getSrcDir() string {
	isModule, isVendor := getInstallationType()
	var srcDir string
	var err error
	if isVendor {
		srcDir, err = filepath.Abs("vendor/github.com/maneac/go-ultralight")
		if err != nil {
			log.Fatalf("Failed to get vendored package directory: %v", err)
		}
	} else {
		srcDir, err = filepath.Abs(os.Getenv("GOPATH"))
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
		fmt.Printf("Downloading %s Ultralight SDK...\n", goos)
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

		if _, err := os.Stat(srcDir); os.IsNotExist(err) {
			fmt.Println("Fetching go-ultralight package...")
			cmd := exec.Command("go", "get", "github.com/maneac/go-ultralight/go-ultralight")
			err := cmd.Run()
			if err != nil {
				log.Fatalf("Failed to get repository: %v", err)
			}
		}
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

func unzip(file, outDir, goos string) {
	fmt.Printf("Extracting %s Ultralight SDK...\n", goos)
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatalf("AAAGH")
	}
	defer r.Close()

	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			os.MkdirAll(filepath.Join(outDir, f.Name), 0777)
		} else if f.Name != ".DS_Store" {
			file, err := f.Open()
			if err != nil {
				log.Fatalf("Failed to open file %s: %v", f.Name, err)
			}
			out, err := os.Create(filepath.Join(outDir, f.Name))
			if err != nil {
				log.Fatalf("Failed to create file %s: %v", f.Name, err)
			}
			io.Copy(out, file)
			file.Close()
			out.Close()
		}
	}
	fmt.Println("SDK extracted!")
}

func untar(file, outDir, goos string) {
	fmt.Printf("Extracting %s Ultralight SDK...\n", goos)
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open tar file: %v", err)
	}
	defer f.Close()

	r, err := xz.NewReader(f, 0)
	if err != nil {
		log.Fatalf("Failed to create tar reader: %v", err)
	}

	tr := tar.NewReader(r)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to get next file in tar: %v", err)
		}

		switch hdr.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(filepath.Join(outDir, hdr.Name), 0777)
			if err != nil {
				log.Fatalf("Failed to make directory %s: %v", hdr.Name, err)
			}
		case tar.TypeReg, tar.TypeRegA:
			w, err := os.Create(filepath.Join(outDir, hdr.Name))
			if err != nil {
				log.Fatalf("Failed to create file %s: %v", hdr.Name, err)
			}
			_, err = io.Copy(w, tr)
			if err != nil {
				log.Fatalf("Failed to copy tar file contents: %v", err)
			}
			w.Close()
		}
	}
	fmt.Println("SDK extracted!")
}

func copyFiles(goos, goarch, srcDir, destDir string) {
	var binaries, libraries []string
	var binPath, libPath string
	var err error
	switch goos {
	case "windows":
		if goarch == "amd64" {
			binPath = filepath.Join(destDir, "SDK", "bin", "win", "x64")
			libPath = filepath.Join(destDir, "SDK", "lib", "win", "x64")
			binaries, err = filepath.Glob(filepath.Join(srcDir, "win", "x64", "*.dll"))
			libraries, err = filepath.Glob(filepath.Join(filepath.Dir(srcDir), "lib", "win", "x64", "*.lib"))
		} else {
			binPath = filepath.Join(destDir, "SDK", "bin", "win", "x86")
			libPath = filepath.Join(destDir, "SDK", "lib", "win", "x86")
			binaries, err = filepath.Glob(filepath.Join(srcDir, "win", "x86", "*.dll"))
			libraries, err = filepath.Glob(filepath.Join(filepath.Dir(srcDir), "lib", "win", "x86", "*.lib"))
		}
	case "darwin":
		binPath = filepath.Join(destDir, "SDK", "bin", "mac")
		binaries, err = filepath.Glob(filepath.Join(srcDir, "*.dylib"))
	case "linux":
		binaries, err = filepath.Glob(filepath.Join(srcDir, "linux", "*.so"))
		binPath = filepath.Join(destDir, "SDK", "bin", "linux")
	}

	if err != nil {
		log.Fatalf("Failed to search package for binaries: %v", err)
	}

	includeBase := filepath.Join(filepath.Dir(srcDir), "include")
	filepath.Walk(includeBase, func(path string, info os.FileInfo, err error) error {
		relPath, err := filepath.Rel(includeBase, path)
		if err != nil {
			log.Fatalf("Failed to get relative path: %v", err)
			return err
		}
		if info.IsDir() {
			os.MkdirAll(filepath.Join(destDir, "SDK", "include", relPath), 0777)
		} else {
			copy(path, filepath.Join(destDir, "SDK", "include", filepath.Dir(relPath)))
		}
		return nil
	})

	if len(binaries) > 0 {
		fmt.Printf("Copying Ultralight binaries to %s...\n", filepath.Dir(binaries[0]))
		if len(libraries) > 0 {
			os.MkdirAll(libPath, 0777)
			for idx := range libraries {
				copy(libraries[idx], libPath)
			}
		}
		os.MkdirAll(binPath, 0777)
		for idx := range binaries {
			copy(binaries[idx], destDir)
			copy(binaries[idx], binPath)
		}
	} else {
		log.Fatalf("No SDK found. Please download and extract the Ultralight SDK as per the instructions found on Github, or run this utility without the 'offline' option.")
	}
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

var binPath, goos, goarch string
var offline, getAll, copyAll, download bool

func main() {
	flag.StringVar(&binPath, "o", ".", "Path to save the binaries to")
	flag.BoolVar(&download, "download", false, "Download SDKs only")
	flag.BoolVar(&offline, "offline", false, "Stops the tool from automatically downloading the Ultralight SDK")
	flag.BoolVar(&getAll, "all", false, "Download the Ultralight SDK for every OS")
	flag.StringVar(&goos, "os", runtime.GOOS, "Target OS to fetch the binaries for")
	flag.StringVar(&goarch, "arch", runtime.GOARCH, "Target OS to fetch the binaries for")
	flag.BoolVar(&copyAll, "targetAll", false, "Fetch the binaries for every OS")
	flag.Parse()

	if binPath == "." && copyAll {
		curDir, err := filepath.Abs(binPath)
		if err != nil {
			log.Fatalf("Failed to get current directory: %v", err)
		}
		binPath = filepath.Join(curDir, "bin")
		err = os.MkdirAll(binPath, 0777)
		if err != nil {
			log.Fatalf("Failed to create binary directory: %v", err)
		}
	} else if binPath != "." {
		curDir, err := filepath.Abs(binPath)
		if err != nil {
			log.Fatalf("Failed to get output directory: %v", err)
		}
		os.MkdirAll(curDir, 0777)
	}

	srcDir := getSrcDir()
	destDir := getDestDir(binPath)

	if !offline {
		if getAll {
			downloadSDK("windows", srcDir)
			downloadSDK("darwin", srcDir)
			downloadSDK("linux", srcDir)
		} else {
			downloadSDK(goos, srcDir)
		}
	}

	if !download {
		srcDir = filepath.Join(srcDir, "SDK", "bin")

		if copyAll {
			copyFiles("windows", "amd64", srcDir, destDir)
			copyFiles("windows", "i386", srcDir, destDir)
			copyFiles("darwin", goarch, srcDir, destDir)
			copyFiles("linux", goarch, srcDir, destDir)
		} else {
			copyFiles(goos, goarch, srcDir, destDir)
		}
		fmt.Printf("Project initialisation complete!\n")
	}
}
