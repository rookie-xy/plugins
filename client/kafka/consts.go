package kafka

import (
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/proxy"
)

const (
    Name = "kafka"
    Namespace = plugin.Flag + "." + proxy.Name + "." + Name
)
