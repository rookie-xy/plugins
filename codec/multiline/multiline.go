package multiline

import (
    "bytes"

    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/register"
)

type Multiline struct {
    log       log.Log
    limit     uint64
    match     byte
    pattern   string
    line     *line
}

func New(l log.Log, v types.Value) (codec.Codec, error) {
    return &Multiline{
        log: l,
    }, nil
}

func (m *Multiline) Encode(in types.Object) (types.Object, error) {
    return nil, nil
}

// Multiline is a split function for a Scanner that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
// The last non-empty line of input will be returned even if it has no
// newline.
func (m *Multiline) Decode(data []byte, atEOF bool) (int, []byte, error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }

    if i := bytes.IndexByte(data, m.match); i >= 0 {
        // We have a full newline-terminated line.
        var multiline []byte

        line := dropCR(data[0:i])
        if match(line) {
            if m.line.Length() > 0 {
                multiline = m.line.Get()
                m.line.Clear()
            }

            m.line.Concat(line)

        } else {
            m.line.Concat(line)
        }

        return i + 1, multiline, nil
    }

    // If we're at EOF, we have a final, non-terminated line. Return it.
    if atEOF {
        return len(data), dropCR(data), nil
    }

    // Request more data.
    return 0, nil, nil
}

func init() {
    register.Codec(Namespace, New)
}
