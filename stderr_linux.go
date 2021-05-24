//+build linux
// https://stackoverflow.com/questions/34772012/capturing-panic-in-golang/34772516

package main

import (
	"log"
	"os"
	"syscall"
)

// redirectStderr to the file passed in
func redirectStderr(f *os.File) {
	err := syscall.Dup3(int(f.Fd()), int(os.Stderr.Fd()), 0)
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}
