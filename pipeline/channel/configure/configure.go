package configure

import (
	"github.com/rookie-xy/hubble/log"
	"github.com/rookie-xy/hubble/types"
	"github.com/mitchellh/mapstructure"
)

type Configure struct {
	log       log.Log

	Max       int
	Duration  string
}

func New(log log.Log) *Configure {
	return &Configure{
        log: log,
        Max: 1024,
        Duration: "10s",
	}
}

func (c *Configure) Init(v types.Value) error {
    if values := v.GetMap(); values != nil {
        if err := mapstructure.Decode(values, c); err != nil {
    		return err
		}
	}
	return nil
}