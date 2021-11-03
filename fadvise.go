// +build !linux

package lumberjack

import (
	"os"
)

func fadvise(_ *os.File,offset,size int) error {
	return nil
}
