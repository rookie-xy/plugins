package log

import (
    "time"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/types/value"
)

type Configure struct {
   *Backoff

    Inactive  time.Duration
    Timeout   time.Duration

    Removed   bool
    Renamed   bool
    EOF       bool
}

func Init(v types.Value, l *Log) error {
	var err error

    vMap := v.GetMap()

    inactive := 3 * time.Second
    if v, ok := vMap["inactive"]; ok {
        inactive, err = time.ParseDuration(v.(string))
        if err != nil {
            return err
        }
    }

    timeout := 10 * time.Second
    if v, ok := vMap["timeout"]; ok {
        timeout, err = time.ParseDuration(v.(string))
        if err != nil {
            return err
        }
    }

    removed := false
    if v, ok := vMap["removed"]; ok {
        removed = v.(bool)
    }

    renamed := false
    if v, ok := vMap["renamed"]; ok {
        removed = v.(bool)
    }

    eof := false
    if v, ok := vMap["eof"]; ok {
        eof = v.(bool)
    }

    if v, ok := vMap["backoff"]; ok {
        this := value.New(v)
        vMap = this.GetMap()
    }


    min := 3 * time.Second
    if v, ok := vMap["min"]; ok {
        min, err = time.ParseDuration(v.(string))
        if err != nil {
            return err
        }
    }

    max := 10 * time.Second
    if v, ok := vMap["max"]; ok {
        timeout, err = time.ParseDuration(v.(string))
        if err != nil {
            return err
        }
    }

    factor := 37
    if v, ok := vMap["factor"]; ok {
        factor = v.(int)
    }

    backoff := &Backoff{
        Min: min,
        Max: max,
        Factor:factor,
    }

    configure := &Configure{
    	Backoff:  backoff,
        Inactive: inactive,
        Timeout:  timeout,
        Removed:  removed,
        Renamed:  renamed,
        EOF:      eof,
    }

    l.conf = configure

    return nil
}
