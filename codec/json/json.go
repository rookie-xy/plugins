package json

import (
    "bytes"
    "encoding/json"

    "github.com/urso/go-structform/gotype"


    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/codec"

	"github.com/rookie-xy/hubble/register"
)

type Json struct {
    log.Log

    name     string
	buffer   bytes.Buffer
	folder  *gotype.Iterator
	pretty   bool
}

func New(l log.Log, v types.Value) (codec.Codec, error) {

    j := &Json{pretty: false}
	if err := reset(j.buffer, j.folder); err != nil {
		panic(err)
	}

    return &Json{
        Log: l,
    }, nil
}

func (j *Json) Encode(in types.Object) (types.Object, error) {
    j.buffer.Reset()

	err := j.folder.Fold(makeEvent(index, e.version, event))
	if err != nil {
		if err := reset(j.buffer, j.folder); err != nil {
			return nil, err
		}

		return nil, err
	}

	data := j.buffer.Bytes()
	if !j.pretty {
		return data, nil
	}

	var buffer bytes.Buffer
	if err = json.Indent(&buffer, data, "", "  "); err != nil {
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
