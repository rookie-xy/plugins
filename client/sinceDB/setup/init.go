package setup

import (
	"os"
	"fmt"
    "path/filepath"

    "github.com/rookie-xy/hubble/paths"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/plugins/client/sinceDB/dump"
)

func Init(v types.Value) (string, error) {
	file := v.GetMap()["path"].(string)

    file = paths.Resolve(paths.Data, file)

    path := filepath.Dir(file)
    err := os.MkdirAll(path, 0750)
    if err != nil {
        return "", fmt.Errorf("Failed to created sinceDB file dir %s: %v", path, err)
    }

    fileInfo, err := os.Lstat(file)
    if os.IsNotExist(err) {
        fmt.Printf("No sinceDB file found under: %s. Creating a new sinceDB file.", file)
        return "", dump.File(file, nil)
	}

    if err != nil {
        return "", err
	}

    if !fileInfo.Mode().IsRegular() {
        if fileInfo.IsDir() {
            return "", fmt.Errorf("SinceDB file path must be a file. %s is a directory.", file)
        }
        return "", fmt.Errorf("SinceDB file path is not a regular file: %s", file)
    }

	fmt.Printf("SinceDB file set to: %s", file)

    return file, nil
}
