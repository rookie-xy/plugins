package yaml

import (
    yml "gopkg.in/yaml.v2"
    "github.com/rookie-xy/worker/src/plugin/codec"
    "github.com/rookie-xy/worker/src/prototype"
    "github.com/rookie-xy/worker/src/register"
)

const Name = "yaml"

type Yaml struct {
    name string
}

func New() *Yaml {
    return &Yaml{}
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
    register.Codec(Name, func(cfg *codec.Config) (codec.Codec, error) {
        return New(), nil
    })
}

/*
func (r *Yaml) Clone() codec.Codec {
    yaml := NewYaml()
    yaml.name = "yaml"
    return yaml
}

func (r *Yaml) Init(configure prototype.Object) int {
    return state.Ignore
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

func (r *Yaml) Type(name string) int {
    fmt.Println("mengshiiiiiiiiiiiiiiiiiiii", name)
    if r.name != name {
        return state.Declined
    }

    fmt.Println("dddddddddddddddddddddddddddddd", name)

    return state.Ok
}

func init() {
    register.Codecs("yml", yaml)
}
*/
