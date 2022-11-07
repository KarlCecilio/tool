package karl_file

import "os"

func PathExist(fullPath string) bool {
	_, err := os.Stat(fullPath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
