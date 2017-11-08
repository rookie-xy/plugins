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
        state.On = true
        state.TTL = -2
        states[index] = state
    }

    return states
}
/*
func (r *sinceDB) ID(value types.Value) string {
    file := value.GetMap()
    return fmt.Sprintf("%d-%d", file["inode"], file["device"])
}

func (s *sinceDB) update(state state.State) error {
    for _, value := range s.values {
    	if s.ID(value) != e.ID() {
    	    continue
        }

        //s.values[i] = e.Value()
    }

    return nil
}
*/
