package states

import (
	"os"
    "encoding/json"

    "github.com/rookie-xy/hubble/models/file"
    "github.com/rookie-xy/hubble/log"
  . "github.com/rookie-xy/hubble/log/level"
)

func Load(path string, log log.Factory) ([]file.State, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    log(DEBUG,"Loading sinceDB data from %s", path)

    decoder := json.NewDecoder(f)
    states  := []file.State{}
    if err = decoder.Decode(&states); err != nil {
        return nil, err
    }

    states = reset(states, log)
    log(DEBUG,"States Loaded from sinceDB: %+v", len(states))

	return states, nil
}

func reset(states []file.State, log log.Factory) []file.State {
    for index, state := range states {
        state.Finished = true
        state.TTL = -2
        states[index] = state
    }

    log(DEBUG,"reset states success")
    return states
}
