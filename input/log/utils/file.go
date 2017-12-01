package utils

import (
	"os"
	"github.com/rookie-xy/hubble/log"
  . "github.com/rookie-xy/hubble/log/level"
)

// IsSameFile checks if the given File path corresponds with the FileInfo given
func SameFile(path string, info os.FileInfo, log log.Factory) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		log(ERROR,"Error during file comparison: %s with %s - Error: %s",
                     path, info.Name(), err)
		return false
	}

	return os.SameFile(fileInfo, info)
}
