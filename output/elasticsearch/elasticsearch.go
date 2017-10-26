package elasticsearch

import (
    "strings"

    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/output"
)

type elasticsearch struct {
    log       log.Log
    pipeline  pipeline.Queue
}

func open(l log.Log, v types.Value) (output.Output, error) {
    elasticsearch := &elasticsearch{
        log: l,
    }

    plugin := pipeline.Plugin

    if value := v.GetMap(); value != nil {
        for key, _ := range value {
            key := key.(string)
            if n := strings.Index(key, "."); n > -1 {
                if key[0:n] == pipeline.Name {
                    plugin = key[n+1 : len(key)]
                }
            }
        }
    }

    if pipeline, err := factory.Pipeline(plugin, l, v); err != nil {
        elasticsearch.pipeline = pipeline
    }

    if queue := factory.Queue(Name); queue != nil {
        if err := queue.Enqueue(elasticsearch.pipeline.(event.Event)); err != nil {
            return nil, err
        }
    }

    return elasticsearch, nil
}

func (r *elasticsearch) Sender(e event.Event) error {
    r.pipeline.Enqueue(e)
    return nil
}

func (r *elasticsearch) Close() int {
    return state.Ok
}

func init() {
    register.Output(Namespace, open)
}
