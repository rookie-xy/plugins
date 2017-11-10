package sinceDB

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/adapter"
    "github.com/rookie-xy/hubble/output"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/models/file"
)

type sinceDB struct {
    adapter.SinceDB
    log       log.Log
    pipeline  pipeline.Queue
}

func open(l log.Log, v types.Value) (output.Output, error) {
    sinceDB := &sinceDB{
        log: l,
    }

    // Open the sinceDB connection channel
    if pipeline := factory.Queue(v.GetString()); pipeline != nil {
        sinceDB.pipeline = pipeline
    }

    // Open the sinceDB client and get the file models
    if sinceDb, err := factory.Forward(plugin.Flag + "." + v.GetString()); err != nil {
        return nil, err
    } else {
        sinceDB.SinceDB = adapter.FileSinceDB(sinceDb)
    }

    return sinceDB, nil
}

func (s *sinceDB) Sender(e event.Event) error {
    s.pipeline.Enqueue(e)
    return nil
}

func (s *sinceDB) Load() []file.State {
    return s.SinceDB.Load()
}

func (s *sinceDB) Close() {
}

func init() {
    register.Output(Namespace, open)
}
