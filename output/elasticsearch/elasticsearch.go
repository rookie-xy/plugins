package elasticsearch

import (
    "strings"

    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/factory"
    "github.com/rookie-xy/hubble/pipeline"
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/adapter"
    "github.com/rookie-xy/hubble/output"
    "github.com/rookie-xy/hubble/proxy"
)

type elasticsearch struct {
    log       log.Log
    pipeline  pipeline.Queue
    value     types.Value
    domain    string
}

func Elasticsearch(l log.Log, v types.Value) (output.Output, error) {
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

    elasticsearch := &elasticsearch{
        log:    l,
        value:  v,
        domain: domain,
    }

    return elasticsearch, nil
}

func (e *elasticsearch) New() (proxy.Forward, error) {
    elasticsearch := &elasticsearch{
    	log: e.log,
    }

    if pipeline, err := factory.Pipeline(e.domain, e.log, e.value); err != nil {
        return nil, err
    } else {
        elasticsearch.pipeline = pipeline
    }

    if queue := factory.Queue(Name); queue != nil {
        if err := queue.Enqueue(adapter.Pipeline2Event(elasticsearch.pipeline)); err != nil {
            return nil, err
        }
    }

    return elasticsearch, nil
}

func (r *elasticsearch) Sender(e event.Event) error {
    return r.pipeline.Enqueue(e)
}

func (r *elasticsearch) Close() {
    r.pipeline.Close()
}

func init() {
    register.Output(Namespace, Elasticsearch)
}

/*
    if pipeline, err := factory.Pipeline(domain, l, v); err != nil {
        return nil, err
    } else {
        elasticsearch.pipeline = pipeline
        fmt.Println("pipelin")
    }

    if queue := factory.Queue(Name); queue != nil {
        if err := queue.Enqueue(adapter.Pipeline2Event(elasticsearch.pipeline)); err != nil {
            return nil, err
        }
    }
*/

