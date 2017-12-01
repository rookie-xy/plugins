package utils

import "os"

func Rotate(path, temp string) error {
    if e := os.Rename(temp, path); e != nil {
        return e
    }

    return nil
}
