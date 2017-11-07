package sinceDB

import (
    "os"
    "fmt"
    "encoding/json"
    "path/filepath"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/paths"
//    "github.com/rookie-xy/hubble/adapter"
    "github.com/rookie-xy/hubble/adapter"
)

type sinceDB struct {
    log     log.Log
    path    string
    states  []adapter.FileState
}

func (r *sinceDB) Init() error {
    r.path = paths.Resolve(paths.Data, r.path)

    path := filepath.Dir(r.path)
    err := os.MkdirAll(path, 0750)
    if err != nil {
        return fmt.Errorf("Failed to created sinceDB file dir %s: %v", path, err)
    }

    fileInfo, err := os.Lstat(r.path)
    if os.IsNotExist(err) {
        fmt.Printf("No sinceDB file found under: %s. Creating a new sinceDB file.", r.path)
        return r.diskDump()
	}

    if err != nil {
        return err
	}

    if !fileInfo.Mode().IsRegular() {
        if fileInfo.IsDir() {
            return fmt.Errorf("SinceDB file path must be a file. %s is a directory.", r.path)
        }
        return fmt.Errorf("SinceDB file path is not a regular file: %s", r.path)
    }

	fmt.Printf("SinceDB file set to: %s", r.path)

    return nil
}

func (r *sinceDB) load() error {
    f, err := os.Open(r.path)
    if err != nil {
        return err
    }
    defer f.Close()

    fmt.Printf("Loading sinceDB data from %s\n", r.path)

    decoder := json.NewDecoder(f)
    states  := []state.State{}
    if err = decoder.Decode(&states); err != nil {
        return fmt.Errorf("Error decoding states: %s", err)
    }

    r.states = reset(states)
    fmt.Printf("States Loaded from sinceDB: %+v\n", len(r.states))

	return nil
}

func reset(states []state.State) []state.State {
    for index, state := range states {
        state.On()
        states[index] = state
//        state.Finished = true
//        state.TTL = -2
    }

    return states
}

func rotate(path, temp string) error {
    if e := os.Rename(temp, path); e != nil {
        fmt.Printf("Rotate error: %s", e)
        return e
    }

    return nil
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    sinceDB := &sinceDB{
        log: l,
    }

    if v != nil {
        sinceDB.path = v.GetMap()["path"].(string)
    }

    if err := sinceDB.Init(); err != nil {
        return nil, err
    }

    if err := sinceDB.load(); err != nil {
    	return nil, fmt.Errorf("Error loading state: %v", err)
	}

    return sinceDB, nil
}

func (r *sinceDB) ID(value types.Value) string {
    file := value.GetMap()
    return fmt.Sprintf("%d-%d", file["inode"], file["device"])
}

func (s *sinceDB) update(state state.State) error {
	/*
    for _, value := range s.values {
    	if s.ID(value) != e.ID() {
    	    continue
        }

        //s.values[i] = e.Value()
    }
	*/

    return nil
}

func (r *sinceDB) diskDump() error {
    fmt.Printf("write sinceDB file: %s\n", r.path)

    temp := r.path + ".new"
    f, err := os.OpenFile(temp, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0600)
    if err != nil {
        fmt.Printf("Failed to create tempfile (%s) for writing: %s", temp, err)
        return err
    }

    encoder := json.NewEncoder(f)
    if err = encoder.Encode(r.states); err != nil {
        f.Close()
        fmt.Printf("Error when encoding the states: %s", err)
        return err
    }

    f.Close()
    err = rotate(r.path, temp)

    fmt.Printf("SinceDB file updated. %d states written.", len(r.states))
    return err
}

func (r *sinceDB) Sender(e event.Event) error {
	fmt.Println("sincedb senderrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
    batch := true
	if batch {

	    /*
        events := adapter.ToEvents(e)
        for _, event := range events.Batch() {
        	if err := r.update(event); err != nil {
        	    return err
            }
        }
	    */

    } else {
        if err := r.update(e); err != nil {
            return err
        }
    }

    return nil
}

func (r *sinceDB) Commit(e event.Event) bool {
    if e != nil {
        r.states = append(r.states, adapter.ToFileState(e))
        return true
    }

    return false
}

func (r *sinceDB) Senders() ([]event.Event, error) {
	var events []event.Event

    for _, state := range r.states {
        events = append(events, state.(event.Event))

        if err := r.update(state); err != nil {
            return events, err
        }
    }

    if err := r.diskDump(); err != nil {
        return events, err
    }

    return nil, nil
}

func (r *sinceDB) Get() []adapter.FileState {
    return r.states
}

func (r *sinceDB) Close() int {
    return state.Ok
}

func init() {
    register.Client(Namespace, open)
}
