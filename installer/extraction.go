package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/xi2/xz"
)

func unzip(file, outDir, goos string) {
	fmt.Printf("Extracting %s Ultralight SDK...\n", goos)
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatalf("Failed to open SDK for extraction: %v", err)
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
