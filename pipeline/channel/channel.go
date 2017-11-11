package channel

import (
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/types"

    "time"
    "go/ast"
    "github.com/rookie-xy/plugins/pipeline/channel/configure"
)

type channel struct {
    log.Log

    channel   chan event.Event
    timer    *time.Ticker
}

func open(l log.Log, v types.Value) (pipeline.Queue, error) {
	configure :=  configure.New(l)
	if err := configure.Init(v); err != nil {
	    return nil, err
    }

    duration, err := time.ParseDuration(configure.Duration)
    if err != nil {
        return nil, err
    }

    return &channel{
        Log: l,
        channel: make(chan event.Event, configure.Max),
        timer: time.NewTicker(duration),
    }, nil
}

// TODO 确定如何保证并发
func (r *channel) Clone() pipeline.Queue {
    return r
}

func (r *channel) Enqueue(e event.Event) error {
    r.channel <- e
    return nil
}

func (r *channel) Dequeue() (event.Event, error) {
    event, closed := <- r.channel
    if closed {
        return event, pipeline.ErrClosed
    }

    return event, nil
}

func (c *channel) Dequeues(size int) ([]event.Event, error) {
    var events []event.Event

    count := 0
    for {
        select {

        case event, closed := <-c.channel:
        	if closed {
        		if count > 0 {
                    return events, pipeline.ErrClosed
                }

                return nil, pipeline.ErrClosed
            }

        	events = append(events, event)
            count++

            if count < size {
                continue
            }

            return events, nil

        case c.timer.C:
        	if count > 0 {
                return events, pipeline.ErrEmpty
            }

            return nil, pipeline.ErrEmpty
        }
    }

    return nil, nil
}

func (r *channel) Requeue(e event.Event) error {
    return r.Enqueue(e)
}

func (r *channel) Close() int {
    return -1
}

func (r *channel) On() bool {
    return true
}

func (r *channel) Off() {
    return
}

func init() {
    register.Pipeline(Namespace, open)
}
