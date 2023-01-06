package dropper

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var ErrNOutOfRange = errNOutOfRange()

func errNOutOfRange() error {
	return errors.New(" dropper: n out of range. n < 1 or n > 3")
}

const DropCacheFile = "/proc/sys/vm/drop_caches"

func NoPermissions() bool {
	_, err := os.OpenFile(DropCacheFile, os.O_WRONLY, 0200)
	return os.IsPermission(err)
}

func Drop(n int) error {
	if n < 1 || n > 3 {
		return ErrNOutOfRange
	}

	err := exec.Command("sync").Run()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(DropCacheFile, os.O_WRONLY, 0200)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprint(n))
	if err != nil {
		return err
	}
	return nil
}
