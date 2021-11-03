// +build !linux

package lumberjack

import (
	"os"
)

func fadvise(_ *os.File,_,_ int) error {
	return nil
}
