package states

import (
	"os"
	"fmt"
    "encoding/json"
    "github.com/rookie-xy/hubble/models/file"
)

func Load(path string) ([]file.State, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    fmt.Printf("Loading sinceDB data from %s\n", path)

    decoder := json.NewDecoder(f)
    states  := []file.State{}
    if err = decoder.Decode(&states); err != nil {
        return nil, fmt.Errorf("Error decoding states: %s", err)
    }

    states = reset(states)
    fmt.Printf("States Loaded from sinceDB: %+v\n", len(states))

	return states, nil
}

func reset(states []file.State) []file.State {
    for index, state := range states {
        state.Finished = true
        state.TTL = -2
        states[index] = state
    }

    return states
}
