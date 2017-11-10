package kafka

import (
	"github.com/rookie-xy/hubble/plugin"
	"github.com/rookie-xy/hubble/output"
)

const (
	Name = "kafka"
	Namespace = plugin.Flag + "." + output.Name + "." + Name
)