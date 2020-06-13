package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func computeFileHash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic("Could not open file " + path + " to compute hash. error: " + err.Error())
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		panic("Computing hash for file  " + path + " failed with error: " + err.Error())
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
