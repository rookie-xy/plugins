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
    "github.com/rookie-xy/hubble/proxy"
)

type kafka struct {
    log       log.Log
    pipeline  pipeline.Queue
    value     types.Value
    domain    string
}

func Kafka(l log.Log, v types.Value) (output.Output, error) {
    domain, ok := plugin.Domain(pipeline.Name, pipeline.Plugin)
    if ok {
        if value := v.GetMap(); value != nil {
            for key, _ := range value {
                key := key.(string)
                if n := strings.Index(key, "."); n > -1 {
                    if key[0:n] == pipeline.Name {
                        domain, _ = plugin.Name(key)
                    }
                }
            }
        }
    }

    kafka := &kafka{
        log:    l,
        value:  v,
        domain: domain,
    }

    return kafka, nil
}

func (k *kafka) New() proxy.Forward {
    kafka := &kafka{
    	log: k.log,
    }

    if pipeline, err := factory.Pipeline(k.domain, k.log, k.value); err != nil {
        return nil
    } else {
        kafka.pipeline = pipeline
    }

    if queue := factory.Queue(Name); queue != nil {
        if err := queue.Enqueue(adapter.Pipeline2Event(kafka.pipeline)); err != nil {
            return nil
        }
    }

    return kafka
}

func (k *kafka) Sender(e event.Event) error {
    return k.pipeline.Enqueue(e)
}

func (k *kafka) Close() {
    k.pipeline.Close()
}

func init() {
    register.Output(Namespace, Kafka)
}
