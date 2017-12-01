package console

import (
    "fmt"

    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/proxy"
    "github.com/rookie-xy/hubble/event"
    "github.com/rookie-xy/hubble/adapter"
    "github.com/rookie-xy/hubble/register"
    "bufio"
    "os"
)

type console struct {
    log.Log
    out     *os.File
 	writer  *bufio.Writer
}

func open(l log.Log, v types.Value) (proxy.Forward, error) {
    console := &console{
        Log: l,
        out: os.Stdout,
    }

    console.writer = bufio.NewWriterSize(console.out, 8*1024)
    return console, nil
}

func (c *console) Sender(e event.Event) error {
    fileEvent := adapter.ToFileEvent(e)
    state := fileEvent.GetFooter()
    body := adapter.ToFileEvent(e).GetBody()
    fmt.Printf("consoleeeeeeeeeeeeeeeeeeeeeeeeeeee: %d#%s\n ", state.Offset, string(body.GetContent()))
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
