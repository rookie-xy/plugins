package configure

import (
    "github.com/rookie-xy/worker/src/configure"
)

type file struct {
    name string
}

func New() *file {
    return &file{}
}

func (r *file) GetFile() {

}

func init() {
    configure.Plugins = append(configure.Plugins, New())
}
