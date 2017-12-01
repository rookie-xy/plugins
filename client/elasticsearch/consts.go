package elasticsearch

import (
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/proxy"
)

const (
    Name = "elasticsearch"
    Namespace = plugin.Flag + "." + proxy.Name + "." + Name
)
