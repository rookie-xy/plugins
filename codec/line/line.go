package line

import (
    "bytes"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/register"
)

type Line struct {
    log    log.Log
    limit  uint64
    match  byte
}

func New(l log.Log, v types.Value) (codec.Codec, error) {
    return &Line{
        log: l,
    }, nil
}

func (l *Line) Encode(in types.Object) (types.Object, error) {
    return nil, nil
}

// Line is a split function for a Scanner that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
// The last non-empty line of input will be returned even if it has no
// newline.
func (l *Line) Decode(data []byte, atEOF bool) (int, []byte, error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }

//    if i := bytes.IndexByte(data, '\n'); i >= 0 {
    if i := bytes.IndexByte(data, l.match); i >= 0 {
        // We have a full newline-terminated line.
        return i + 1, dropCR(data[0:i]), nil
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
