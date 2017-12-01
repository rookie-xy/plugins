package path

import (
	"os"
	"fmt"
    "errors"
    "path/filepath"

    "github.com/rookie-xy/hubble/paths"
    "github.com/rookie-xy/hubble/types"
	"github.com/rookie-xy/hubble/models/file"

    "github.com/rookie-xy/plugins/client/sinceDB/dump"
	"github.com/rookie-xy/hubble/log"
  . "github.com/rookie-xy/hubble/log/level"
)

func Init(v types.Value, states *file.States, log log.Factory) (string, error) {
	var file string

    if values := v.GetMap(); values != nil {
    	if value, ok := values["file"]; ok {
            file = value.(string)
		} else {
            return "", errors.New("not found sinceDB file")
		}
	}

    file = paths.Resolve(paths.Data, file)

    path := filepath.Dir(file)
    err := os.MkdirAll(path, 0750)
    if err != nil {
        return file, fmt.Errorf("failed to created sinceDB file dir %s: %v", path, err)
    }

    fileInfo, err := os.Lstat(file)
    if os.IsNotExist(err) {
        log(WARN,"No sinceDB file found under: %s. Creating a new sinceDB file", file)
        return file, dump.File(file, states, log)
	}

    if err != nil {
        return file, err
	}

    if !fileInfo.Mode().IsRegular() {
        if fileInfo.IsDir() {
            return "", fmt.Errorf("SinceDB file path must be a file. %s is a directory", file)
        }
        return file, fmt.Errorf("SinceDB file path is not a regular file: %s", file)
    }

	log(DEBUG,"SinceDB file set to: %s\n", file)
    return file, nil
}
