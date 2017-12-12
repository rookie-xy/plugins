package dump

import (
	"os"
	"encoding/json"

    "github.com/rookie-xy/plugins/client/sinceDB/utils"
    "github.com/rookie-xy/hubble/models/file"
    "github.com/rookie-xy/hubble/log"
  . "github.com/rookie-xy/hubble/log/level"
)

func File(path string, states *file.States, log log.Factory) error {
    log(DEBUG,"write states to file: %s\n", path)

    temp := path + ".new"
    f, err := os.OpenFile(temp, flag, model)
    if err != nil {
        log(ERROR,"Failed to create temp file (%s) for writing: %s", temp, err)
        return err
    }

    this := states.Get()
    encoder := json.NewEncoder(f)
    if err = encoder.Encode(this); err != nil {
        f.Close()
        log(ERROR,"Error when encoding the states: %s", err)
        return err
    }

    f.Close()
    err = utils.Rotate(path, temp)

    log(DEBUG,"SinceDB file updated. %d states written.", len(this))
    return err
}
