package elasticsearch

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/state"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/output"
)

const Namespace = "plugin.output.elasticsearch"

type elasticsearch struct {
    log       log.Log
    pipeline  pipeline.Queue
}

func open(l log.Log, v types.Value) (output.Output, error) {
    elasticsearch := &elasticsearch{
        log: l,
    }

    if pipeline := factory.Queue(v.GetString()); pipeline != nil {
        elasticsearch.pipeline = pipeline
    }

    return elasticsearch, nil
}

func (r *elasticsearch) Sender(e event.Event, batch bool) error {
    r.pipeline.Enqueue(e)

    return nil
}

func (r *elasticsearch) Close() int {
    return state.Ok
}

func init() {
    register.Output(Namespace, open)
}
