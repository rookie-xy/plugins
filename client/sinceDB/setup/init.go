package setup

import (
	"os"
	"fmt"
    "errors"
    "path/filepath"

    "github.com/rookie-xy/hubble/paths"
    "github.com/rookie-xy/hubble/types"
	"github.com/rookie-xy/hubble/models/file"

    "github.com/rookie-xy/plugins/client/sinceDB/dump"
)

func Init(v types.Value, states *file.States) (string, error) {
	var file string

    if values := v.GetMap(); values != nil {
    	if value, ok := values["file"]; ok {
            file = value.(string)
		} else {
            return "", errors.New("Not found sinceDB file")
		}
	}

    file = paths.Resolve(paths.Data, file)

    path := filepath.Dir(file)
    err := os.MkdirAll(path, 0750)
    if err != nil {
        return file, fmt.Errorf("Failed to created sinceDB file dir %s: %v", path, err)
    }

    fileInfo, err := os.Lstat(file)
    if os.IsNotExist(err) {
        fmt.Printf("No sinceDB file found under: %s. Creating a new sinceDB file\n", file)
        return file, dump.File(file, states)
	}

    if err != nil {
        return file, err
	}

    if !fileInfo.Mode().IsRegular() {
        if fileInfo.IsDir() {
            return "", fmt.Errorf("SinceDB file path must be a file. %s is a directory.", file)
        }
        return file, fmt.Errorf("SinceDB file path is not a regular file: %s", file)
    }

	fmt.Printf("SinceDB file set to: %s\n", file)

    return file, nil
}
