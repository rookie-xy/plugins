package channel

import (
    "github.com/rookie-xy/hubble/src/event"
    "github.com/rookie-xy/hubble/src/state"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/register"
    "github.com/rookie-xy/hubble/src/pipeline"
    "github.com/rookie-xy/hubble/src/types"

)

const Namespace = "plugin.pipeline.channel"

type channel struct {
    log.Log
    channel chan event.Event
}

func open(l log.Log, v types.Value) (pipeline.Queue, error) {
    return &channel{
        Log: l,
        channel: make(chan event.Event, 1024),
    }, nil
}

// TODO 确定如何保证并发
func (r *channel) Clone() pipeline.Queue {
    return r
}

func (r *channel) Enqueue(e event.Event) int {
    r.channel <- e
    return state.Ok
}

func (r *channel) Dequeue(size int) (event.Event, int) {
    event, open := <- r.channel
    if !open {
        return nil, state.Done
    }

    return event, state.Ok
}

func (r *channel) Requeue(size int) (event.Event, int) {
    return r.Dequeue(size)
}

func (r *channel) Close() int {
    return state.Ok
}

func init() {
    register.Pipeline(Namespace, open)
}
