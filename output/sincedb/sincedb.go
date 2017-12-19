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
    "errors"
    "github.com/rookie-xy/hubble/proxy"
)

type sinceDB struct {
    adapter.SinceDB
    log       log.Log
    pipeline  pipeline.Queue
}

func SinceDB(l log.Log, v types.Value) (output.Output, error) {
    sinceDB := &sinceDB{
        log: l,
    }

    // Open the sinceDB connection channel
    if pipeline := factory.Queue(v.GetString()); pipeline != nil {
        sinceDB.pipeline = pipeline
    }

    // Open the sinceDB client and get the file models
    if key, ok := plugin.Name(v.GetString()); ok {
        if sinceDb, err := factory.Forward(key); err != nil {
            return nil, err
        } else {
            sinceDB.SinceDB = adapter.FileSinceDB(sinceDb)
        }
    } else {
        return nil, errors.New("plugin name error")
    }
    return sinceDB, nil
}

func (s *sinceDB) New() (proxy.Forward, error) {
    return &sinceDB{
        log:      s.log,
        SinceDB:  s.SinceDB,
        pipeline: s.pipeline,
    }, nil
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
    register.Output(Namespace, SinceDB)
}
