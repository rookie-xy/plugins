package json

import (
    "bytes"
    "encoding/json"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/codec"

	"github.com/rookie-xy/hubble/register"
	"github.com/mitchellh/mapstructure"
	"github.com/rookie-xy/hubble/adapter"
	"github.com/rookie-xy/hubble/event"
)

type Json struct {
    log.Log

	buffer   bytes.Buffer
	Pretty   bool
}

func New(l log.Log, v types.Value) (codec.Codec, error) {
    j := &Json{
    	Log: l,
    	Pretty: false,
    }

    if err := mapstructure.Decode(v.GetMap(), j); err != nil {
        return nil, err
    }

    return j, nil
}

func (j *Json) Encode(in types.Object) ([]byte, error) {
    data, err := json.Marshal(adapter.ToFileEvent(in.(event.Event)))
    if err != nil {
    	return nil, err
	}

	if !j.Pretty {
		return data, nil
	}

	var buffer bytes.Buffer
	if err := json.Indent(&buffer, data, "", "  "); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (j *Json) Decode(in []byte, atEOF bool) (int, []byte, error) {
    return 0, nil, nil
}

func init() {
    register.Codec(Namespace, New)
}
