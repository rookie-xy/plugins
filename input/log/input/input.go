package input

import (
    "time"
    "github.com/rookie-xy/hubble/types"
    "github.com/mitchellh/mapstructure"
)

type Input struct {
    Inactive  time.Duration
    Timeout   time.Duration

    Removed   bool
    Renamed   bool
    EOF       bool

    Backoff   Backoff
}

func New() *Input {
   return &Input {
       Inactive: 3 * time.Second,
       Timeout: 10 * time.Second,
       Removed: false,
       Renamed: false,
       EOF: false,
       Backoff: Backoff{
           Min: 3 *  time.Second,
           Max: 10 * time.Second,
           Factor: 37,
       },
    }
}

func (i *Input) Init(v types.Value) error {
    if values := v.GetMap(); values != nil {
    	if err := mapstructure.Decode(values, i); err != nil {
    		return err
		}
	}

	return nil
}
