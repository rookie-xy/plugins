package slot

import (
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/state"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/register"
    "github.com/rookie-xy/hubble/src/pipeline"
    "github.com/rookie-xy/hubble/src/types"

)

const Namespace = "plugin.pipeline.slot"

type slot struct {
    log.Log
    channel chan event.Event
}

func open(l log.Log, v types.Value) (pipeline.Pipeline, error) {
    return &slot{
        Log: l,
        channel: make(chan event.Event, 1024),
    }, nil
}

// TODO 确定如何保证并发
func (r *slot) Clone() pipeline.Pipeline {
    return r
}

func (r *slot) Push(e event.Event) int {
    r.channel <- e
    return state.Ok
}

func (r *slot) Pull(size int) (event.Event, int) {
    event, open := <- r.channel
    if !open {
        return nil, state.Done
    }

    return event, state.Ok
}

func (r *slot) Close() int {
    return state.Ok
}

func init() {
    register.Pipeline(Namespace, open)
}
