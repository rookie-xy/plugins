package pipeline

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/pipeline"
)

const Namespace = "plugin.client.pipeline.elasticsearch"

type elasticsearch struct {
    log       log.Log
    pipeline  pipeline.Queue
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    elasticsearch := &elasticsearch{
        log: l,
    }

    if clone := factory.Clone(v.GetString()); clone != nil {
        elasticsearch.pipeline = clone.Clone()
    }

    return elasticsearch, nil
}

func (r *elasticsearch) Sender(e event.Event) int {
    r.pipeline.Enqueue(e)

    return state.Ok
}

func (r *elasticsearch) Close() int {
    return state.Ok
}

func init() {
    register.Client(Namespace, open)
}
