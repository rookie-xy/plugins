package pipeline

import (
    "github.com/rookie-xy/worker/src/event"
    "github.com/rookie-xy/worker/src/state"
    "github.com/rookie-xy/worker/src/log"
    "github.com/rookie-xy/worker/src/register"
)

const Namespace = "plugin.channel.pipeline"

type pipeline struct {
    log.Log
    channel chan event.Event
}

func open(log log.Log, size int) *pipeline {
    return &pipeline{
        Log: log,
        channel: make(chan event.Event, size),
    }
}

// TODO 确定如何保证并发
func (r *pipeline) Clone() *pipeline {
    return r
}

func (r *pipeline) Push(e event.Event) int {
    r.channel <- e
    return state.Ok
}

func (r *pipeline) Pull(size int) (event.Event, int) {
    event, open := <- r.channel
    if !open {
        return nil, state.Done
    }

    return event, state.Ok
}

func (r *pipeline) Close() int {
    return state.Ok
}

func init() {
    register.Channel(Namespace, open)
}
