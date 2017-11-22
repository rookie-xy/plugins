package configure

import (
    "time"
    "github.com/rookie-xy/hubble/types"
    "github.com/mitchellh/mapstructure"
)

type Configure struct {
    Inactive  string
    Timeout   string

    inactive  time.Duration
    timeout   time.Duration

    Removed   bool
    Renamed   bool
    EOF       bool

    Backoff   Backoff
}

func New() *Configure {
   return &Configure {
       inactive: 3 * time.Second,
       timeout: 10 * time.Second,
       Removed: false,
       Renamed: false,
       EOF: false,
       Backoff: Backoff{
           min: 3 *  time.Second,
           max: 10 * time.Second,
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

	var err error
    if c.inactive, err = time.ParseDuration(c.Inactive); err != nil {
		return err
	}

    if c.timeout, err = time.ParseDuration(c.Timeout); err != nil {
		return err
	}

    if c.Backoff.max, err = time.ParseDuration(c.Backoff.Max); err != nil {
		return err
	}

    if c.Backoff.min, err = time.ParseDuration(c.Backoff.Min); err != nil {
		return err
	}

	return nil
}

func (c *Configure) GetInactive() time.Duration {
    return c.inactive
}

func (c *Configure) GetTimeout() time.Duration {
    return c.timeout
}
