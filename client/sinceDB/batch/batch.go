package batch

import (
	"github.com/rookie-xy/hubble/types"
	"github.com/rookie-xy/hubble/log"
	"github.com/mitchellh/mapstructure"
	"time"
)

type Batch struct {
    log       log.Log

    Batch     batch
    duration  time.Duration
}

type batch struct {
    Max      int
    Timeout  string
}

func New(log log.Log) *Batch {
	return &Batch{
		log: log,
		Batch: batch{
            Max: 128,
            Timeout: "10s",
		},
	}
}

func (b *Batch) Init(v types.Value) error {
    if values := v.GetMap(); values != nil {
    	if err := mapstructure.Decode(values, b); err != nil {
    		return err
		}
	}

    duration, err := time.ParseDuration(b.Batch.Timeout)
    if err != nil {
        return err
    }

    b.duration = duration
	return nil
}