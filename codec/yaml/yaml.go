package yaml

import (
    yml "gopkg.in/yaml.v2"
    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/register"
    "github.com/rookie-xy/hubble/log"
)

type Yaml struct {
    log.Log
    name string
}

func New(l log.Log, v types.Value) (codec.Codec, error) {
    return &Yaml{
        Log: l,
    }, nil
}

func (r *Yaml) Encode(in types.Object) (types.Object, error) {
    out, error := yml.Marshal(in);
    if error != nil {
        return nil, error
    }

    return out, nil
}

func (r *Yaml) Decode(in []byte, atEOF bool) (int, types.Object, error) {
    var out interface{}

    if e := yml.Unmarshal(in, &out); e != nil {
        return 0, nil, e
    }

    return 0, out, nil
}

func init() {
    register.Codec(Namespace, New)
}
