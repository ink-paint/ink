package util

import "os"

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		existed = false
	}
	return existed
}
