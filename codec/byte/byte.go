package byte

import (
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/register"
)

type Byte struct {
    log  log.Log
}

func New(l log.Log, _ types.Value) (codec.Codec, error) {
    return &Byte{
        log: l,
    }, nil
}

func (b *Byte) Encode(in types.Object) ([]byte, error) {
    return nil, nil
}

// Byte is a split function for a Scanner that returns each byte as a token.
func (b *Byte) Decode(data []byte, atEOF bool) (int, []byte, error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }

    return 1, data[0:1], nil
}

func init() {
    register.Codec(Namespace, New)
}
