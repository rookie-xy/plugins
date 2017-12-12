package console

import (
	"fmt"
    "os"
    "bufio"

	"github.com/mitchellh/mapstructure"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/register"
	"github.com/rookie-xy/hubble/codec"
	"github.com/rookie-xy/hubble/plugin"
	"github.com/rookie-xy/hubble/factory"
	"github.com/rookie-xy/hubble/types/value"
)

type console struct {
    log.Log

    out      *os.File
 	writer   *bufio.Writer
 	Bufsize   int
 	Split     string

 	end       byte
 	codec     codec.Codec
}

func open(log log.Log, v types.Value) (proxy.Forward, error) {
    console := &console{
        Log: log,
        out: os.Stdout,
        Bufsize: 8*512,
        end: '\n',
    }

    if values := v.GetMap(); values != nil {
        if err := mapstructure.Decode(values, console); err != nil {
    		return nil, err
		}

		if len(console.Split) > 0 {
			console.end = []byte(console.Split)[0]
		}

		for k, v := range values {
			if key, ok := plugin.Check(codec.Name, k.(string)); ok {
				if name, ok := plugin.Name(key); ok {
					var err error
					if console.codec, err = factory.Codec(name, log, value.New(v)); err != nil {
						return nil, err
					}
				}
			}
		}
	}

    console.writer = bufio.NewWriterSize(console.out, console.Bufsize)
    return console, nil
}

func (c *console) Sender(e event.Event) error {
   	serializedEvent, err := c.codec.Encode(e)
   	if err != nil {
   		return err
	}

	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaa")

    if err := flush(c.writer, serializedEvent, c.end); err != nil {
		return err
	}

    return nil
}

func (c *console) Close() {
}

func init() {
    register.Client(Namespace, open)
}

    //fileEvent := adapter.ToFileEvent(e)
    //state := fileEvent.GetFooter()
    //adapter.ToFileEvent(e)
    //body := adapter.ToFileEvent(e).GetBody()
    //fmt.Printf("consoleeeeeeeeeeeeeeeeeeeeeeeeeeee: %d#%s\n ", state.Offset, string(body.GetContent()))
