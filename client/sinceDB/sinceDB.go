package sinceDB

import (
    "fmt"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/plugins/client/sinceDB/dump"
    "github.com/rookie-xy/plugins/client/sinceDB/setup"
    "github.com/rookie-xy/plugins/client/sinceDB/states"
    "github.com/rookie-xy/hubble/models/file"
    "github.com/rookie-xy/hubble/adapter"
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

    if v == nil {
        return nil, fmt.Errorf("Error value is nil")
    }

    if path, err := setup.Init(v, sinceDB.states); err != nil {
        return nil, err

    } else {
    	sinceDB.path = path

        if states, err := states.Load(path); err != nil {
            return nil, fmt.Errorf("Error loading states: %v", err)
	    } else {
	        sinceDB.states.Set(states)
        }
    }

    return sinceDB, nil
}

func (r *sinceDB) Sender(e event.Event) error {
	fileEvent := adapter.ToFileEvent(e)
    r.states.Update(fileEvent.GetState())

    if err := dump.File(r.path, r.states); err != nil {
        return err
    }

    return nil
}

func (r *sinceDB) Commit(e event.Event) bool {
    if e != nil {
        if len(r.events) > 2 {
            return false
        }

        r.events = append(r.events, e)
        return true
    }

    return false
}

func (r *sinceDB) Senders() ([]event.Event, error) {
    for _, event := range r.events {
    	fileEvent := adapter.ToFileEvent(event)
    	r.states.Update(fileEvent.GetState())
    }

    if err := dump.File(r.path, r.states); err != nil {
        return r.events, err
    }

    return nil, nil
}

func (r *sinceDB) Load() []file.State {
    return r.states.States
}

func (r *sinceDB) Close() {
}

func init() {
    register.Client(Namespace, open)
}
