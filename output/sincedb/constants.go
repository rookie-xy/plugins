package sinceDB

import (
	"github.com/rookie-xy/hubble/plugin"
	"github.com/rookie-xy/hubble/output"
	"github.com/rookie-xy/hubble/proxy"
)

const (
	Name = "sinceDB"
	Namespace = plugin.Flag + "." + output.Name + "." + Name
	SinceDB = plugin.Flag + "." + proxy.Name + "." + Name
)