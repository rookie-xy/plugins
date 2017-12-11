package console

import (
    "fmt"
    "bufio"
    "os"

	"github.com/mitchellh/mapstructure"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/adapter"
    "github.com/rookie-xy/hubble/register"
	"github.com/rookie-xy/hubble/codec"
)

type console struct {
    log.Log

    out      *os.File
 	writer   *bufio.Writer
 	Bufsize   int
 	End       byte
 	codec     codec.Codec
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    console := &console{
        Log: l,
        out: os.Stdout,
        Bufsize: 8*512,
        End: '\n',
    }

    if values := v.GetMap(); values != nil {
        if err := mapstructure.Decode(values, console); err != nil {
    		return nil, err
		}
	}

    console.writer = bufio.NewWriterSize(console.out, console.Bufsize)
    return console, nil
}

func (c *console) Sender(e event.Event) error {
    fileEvent := adapter.ToFileEvent(e)
    //state := fileEvent.GetFooter()
    //adapter.ToFileEvent(e)
    //body := adapter.ToFileEvent(e).GetBody()
    //fmt.Printf("consoleeeeeeeeeeeeeeeeeeeeeeeeeeee: %d#%s\n ", state.Offset, string(body.GetContent()))
   	serializedEvent, err := c.codec.Encode(fileEvent)
   	if err != nil {
   		return err
	}

	if err := c.writeBuffer(serializedEvent); err != nil {
		//logp.Critical("Unable to publish events to console: %v", err)
		return err
	}

	if err := c.writeBuffer(c.End); err != nil {
		//logp.Critical("Error when appending newline to event: %v", err)
		return err
	}

    return nil
}

func (c *console) Close() {
}

func (c *console) writeBuffer(buf []byte) error {
	written := 0
	for written < len(buf) {
		n, err := c.writer.Write(buf[written:])
		if err != nil {
			return err
		}

		written += n
	}
	return nil
}

func init() {
    register.Client(Namespace, open)
}
