package yaml

import (
    yml "gopkg.in/yaml.v2"
    "github.com/rookie-xy/hubble/src/codec"
    "github.com/rookie-xy/hubble/src/prototype"
    "github.com/rookie-xy/hubble/src/register"
    "github.com/rookie-xy/hubble/src/log"
    "github.com/rookie-xy/hubble/src/command"
)

const Namespace = "plugin.codec.yaml"

type Yaml struct {
    log.Log
    name string
}

func New(l log.Log, c *command.Command) (codec.Codec, error) {
    return &Yaml{
        Log: l,
    }, nil
}

func (r *Yaml) Encode(in prototype.Object) (prototype.Object, error) {
    out, error := yml.Marshal(in);
    if error != nil {
        return nil, error
    }

    return out, nil
}

func (r *Yaml) Decode(in []byte) (prototype.Object, error) {
    var out interface{}

    if e := yml.Unmarshal(in, &out); e != nil {
        return nil, e
    }

    return out, nil
}

func init() {
    register.Codec(Namespace, New)
}
