package sinceDB

import (
    "fmt"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/plugins/client/sinceDB/dump"
    "github.com/rookie-xy/plugins/client/sinceDB/path"
    "github.com/rookie-xy/plugins/client/sinceDB/states"
    "github.com/rookie-xy/hubble/models/file"
    "github.com/rookie-xy/hubble/adapter"
    "github.com/rookie-xy/plugins/client/sinceDB/batch"
)

type sinceDB struct {
    log      log.Log

    path     string
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

    if path, err := path.Init(v, sinceDB.states); err != nil {
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

func (s *sinceDB) Sender(e event.Event) error {
	fileEvent := adapter.ToFileEvent(e)
    s.states.Update(fileEvent.GetFooter())

    if err := dump.File(s.path, s.states); err != nil {
        return err
    }

    return nil
}

func (s *sinceDB) Senders(events []event.Event) error {
    for _, event := range events {
    	fileEvent := adapter.ToFileEvent(event)
    	s.states.Update(fileEvent.GetFooter())
    }

    if err := dump.File(s.path, s.states); err != nil {
        return err
    }

    return nil
}

func (s *sinceDB) Load() []file.State {
    return s.states.States
}

func (s *sinceDB) Close() {
}

func init() {
    register.Client(Namespace, open)
}
