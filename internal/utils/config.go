package utils

import "os"

func ValidatePath(path string) error {
	if _, err := os.Stat(path); err != nil && err == os.ErrNotExist {
		return err
	}
	return nil
}
