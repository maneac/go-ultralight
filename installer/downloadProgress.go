package main

import (
	"fmt"
	"strings"

	"github.com/dustin/go-humanize"
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
