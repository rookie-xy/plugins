package utils

import (
	"os"
	"fmt"
)

func Rotate(path, temp string) error {
    if e := os.Rename(temp, path); e != nil {
        fmt.Printf("Rotate error: %s", e)
        return e
    }

    return nil
}
