package codec


import (
    yml "gopkg.in/yaml.v2"
    "github.com/rookie-xy-old/hubble/hubble/src/prototype"
)

type Yaml struct {
    name string
}

func NewYaml() *Yaml {
    return &Yaml{}
}

var yaml = &Yaml{
    name: "yaml",
}

func (r *Yaml) New() Codec {
    yaml := NewYaml()

    yaml.name = "yaml"

    return yaml
}

func (r *Yaml) Init(configure prototype.Object) int {
    return Ignore
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
    if r.name != name {
        return Ignore
    }

    return Ok
}

func init() {
    Codecs = append(Codecs, yaml)
}
