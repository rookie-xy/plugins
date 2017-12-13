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

func New(l log.Log, _ types.Value) (codec.Decoder, error) {
    return &Byte{
        log: l,
    }, nil
}

func (b *Byte) Decode(in []byte) (types.Object, error) {
    return nil, nil
}

// Byte is a split function for a Scanner that returns each byte as a token.
func (b *Byte) LogDecode(data []byte, atEOF bool) (int, []byte, error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }

    return 1, data[0:1], nil
}

func init() {
    register.Decoder(Namespace, New)
}
