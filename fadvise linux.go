//go:build !linux
// +build !linux

package lumberjack

import (
	"os"
	"golang.org/x/sys/unix"
)

func fadvise(file *os.File,offset,size int) error {
	return unix.Fadvise(int(file.Fd(),offset,size,unix.FADV_NONTNEED))
}
