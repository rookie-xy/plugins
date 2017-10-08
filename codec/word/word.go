package word

import (
    "unicode/utf8"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/register"
)

type Word struct {
    log    log.Log
}

func New(l log.Log, v types.Value) (codec.Codec, error) {
    return &Word{
        log: l,
    }, nil
}

func (w *Word) Encode(in types.Object) (types.Object, error) {
    return nil, nil
}

// ScanWords is a split function for a Scanner that returns each
// space-separated word of text, with surrounding spaces deleted. It will
// never return an empty string. The definition of space is set by
// unicode.IsSpace.
func (w *Word) Decode(data []byte, atEOF bool) (int, []byte, error) {
    // Skip leading spaces.
    start := 0
    for width := 0; start < len(data); start += width {
        var r rune
        r, width = utf8.DecodeRune(data[start:])
        if !isSpace(r) {
            break
        }
    }

    // Scan until space, marking end of word.
    for width, i := 0, start; i < len(data); i += width {
        var r rune
        r, width = utf8.DecodeRune(data[i:])
        if isSpace(r) {
            return i + width, data[start:i], nil
        }
    }

    // If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
    if atEOF && len(data) > start {
        return len(data), data[start:], nil
    }

    // Request more data.
    return start, nil, nil
}

func init() {
    register.Codec(Namespace, New)
}
