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
    "github.com/rookie-xy/plugins/client/sinceDB/models"
    "github.com/rookie-xy/plugins/client/sinceDB/dump"
    "github.com/rookie-xy/plugins/client/sinceDB/setup"
    "github.com/rookie-xy/plugins/client/sinceDB/states"
    "github.com/rookie-xy/hubble/models/file"
)

type sinceDB struct {
    log      log.Log
    path     string
    events   []event.Event
    states  *file.States
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    sinceDB := &sinceDB{
        log: l,
        states: file.News(),
    }

    if path, err := setup.Init(v); err != nil {
        return nil, err

    } else {
    	sinceDB.path = path

        if states, err := states.Load(path); err != nil {
            return nil, fmt.Errorf("Error loading models: %v", err)
	    } else {
	        sinceDB.states.SetStates(states)
        }
    }

    return sinceDB, nil
}

func (r *sinceDB) Sender(e event.Event) error {
    if err := r.states.Update(e); err != nil {
        return err
    }

    if err := dump.File(r.path, r.states); err != nil {
        return err
    }

    return nil
}

func (r *sinceDB) Commit(e event.Event) bool {
    if e != nil {
        r.events = append(r.events, e)
        return true
    }

    return false
}

func (r *sinceDB) Senders() ([]event.Event, error) {
    for index, event := range r.events {
        if err := r.states.Update(event.Getstate()); err != nil {
            r.events = r.events[index:]
            return r.events, err
        }
    }

    if err := dump.File(r.path, r.states); err != nil {
        return r.events, err
    }

    return nil, nil
}

func (r *sinceDB) Load() []file.State {
    return r.states.States
}

func (r *sinceDB) Close() int {
    return state.Ok
}

func init() {
    register.Client(Namespace, open)
}
