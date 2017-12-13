package redis

import (
    "github.com/rookie-xy/hubble/plugin"
    "github.com/rookie-xy/hubble/proxy"
)

const (
    Name = "redis"
    Namespace = plugin.Flag + "." + proxy.Name + "." + Name
)
