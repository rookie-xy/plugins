package sinceDB

import (
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
    "errors"
  . "github.com/rookie-xy/hubble/log/level"
)

type sinceDB struct {
    log.Log
    level    Level

    path     string
    states  *file.States
}

func open(log log.Log, v types.Value) (proxy.Forward, error) {
    sinceDB := &sinceDB{
        Log: log,
        level: adapter.ToLevelLog(log).Get(),
    }
    sinceDB.states = file.News(sinceDB.log)

    if v == nil {
        return nil, errors.New("error value is nil")
    }

    if path, err := path.Init(v, sinceDB.states, sinceDB.log); err != nil {
        return nil, err

    } else {
    	sinceDB.path = path

        if states, err := states.Load(path, sinceDB.log); err != nil {
            return nil, err
	    } else {
	        sinceDB.states.Set(states)
        }
    }
    return sinceDB, nil
}
/*
func (s *sinceDB) Clone() types.Object {
    return &sinceDB{
        Log:    s.Log,
        level:  s.level,
        path:   s.path,
        states: s.states,
    }
}
*/

func (s *sinceDB) Sender(e event.Event) error {
	fileEvent := adapter.ToFileEvent(e)
    s.states.Update(fileEvent.GetFooter())

    if err := dump.File(s.path, s.states, s.log); err != nil {
        return err
    }
    return nil
}

func (s *sinceDB) Senders(events []event.Event) error {
    for _, event := range events {
    	if event != nil {
            fileEvent := adapter.ToFileEvent(event)
            s.states.Update(fileEvent.GetFooter())
        }
    }

    if err := dump.File(s.path, s.states, s.log); err != nil {
        return err
    }
    return nil
}

func (s *sinceDB) Load() []file.State {
    return s.states.States
}

func (s *sinceDB) Close() {
}

func (s *sinceDB) log(l Level, fmt string, args ...interface{}) {
    log.Print(s.Log, s.level, l, fmt, args...)
}

func init() {
    register.Client(Namespace, open)
}
