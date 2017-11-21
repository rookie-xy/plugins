package configure

import (
    "time"
    "github.com/rookie-xy/hubble/types"
    "github.com/mitchellh/mapstructure"
)

type Configure struct {
    Inactive  time.Duration
    Timeout   time.Duration

    Removed   bool
    Renamed   bool
    EOF       bool

    Backoff   Backoff
}

func New() *Configure {
   return &Configure {
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

func (c *Configure) Init(v types.Value) error {
    if values := v.GetMap(); values != nil {
    	if err := mapstructure.Decode(values, c); err != nil {
    		return err
		}
	}

	return nil
}
