package dump

import (
	"fmt"
	"os"
	"encoding/json"

    "github.com/rookie-xy/plugins/client/sinceDB/models"
    "github.com/rookie-xy/plugins/client/sinceDB/utils"
    "github.com/rookie-xy/hubble/models/file"
)

func File(path string, states *file.States) error {
    fmt.Printf("write states to file: %s\n", path)

    temp := path + ".new"
    f, err := os.OpenFile(temp, flag, model)
    if err != nil {
        fmt.Printf("Failed to create tempfile (%s) for writing: %s", temp, err)
        return err
    }

    encoder := json.NewEncoder(f)
    if err = encoder.Encode(states); err != nil {
        f.Close()
        fmt.Printf("Error when encoding the states: %s", err)
        return err
    }

    f.Close()
    err = utils.Rotate(path, temp)

    fmt.Printf("SinceDB file updated. %d states written.", len(r.states))
    return err
}
