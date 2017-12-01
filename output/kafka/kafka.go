package kafka

import (
    "strings"

    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/output"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/adapter"
)

type kafka struct {
    log       log.Log
    pipeline  pipeline.Queue
}

func open(l log.Log, v types.Value) (output.Output, error) {
    kafka := &kafka{
        log: l,
    }

    pluginName := plugin.Flag + "." + pipeline.Name + "." + pipeline.Plugin

    if value := v.GetMap(); value != nil {
        for key, _ := range value {
            key := key.(string)
            if n := strings.Index(key, "."); n > -1 {
                if key[0:n] == pipeline.Name {
                    pluginName = plugin.Flag + "." + key
                }
            }
        }
    }

    if pipeline, err := factory.Pipeline(pluginName, l, v); err != nil {
        return nil, err
    } else {
        kafka.pipeline = pipeline
    }

    if queue := factory.Queue(Name); queue != nil {
        if err := queue.Enqueue(adapter.Pipeline2Event(kafka.pipeline)); err != nil {
            return nil, err
        }
    }

    return kafka, nil
}

func (k *kafka) Sender(e event.Event) error {
    return k.pipeline.Enqueue(e)
}

func (k *kafka) Close() {
    k.pipeline.Close()
}

func init() {
    register.Output(Namespace, open)
}
