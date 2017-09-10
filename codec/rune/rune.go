package rune

import (
    "unicode/utf8"
    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/register"
)

type Rune struct {
    log    log.Log
}

func New(l log.Log, v types.Value) (codec.Codec, error) {
    return &Rune{
        log: l,
    }, nil
}

func (r *Rune) Encode(in types.Object) (types.Object, error) {
    return nil, nil
}

// Rune decode is a split function for a Scanner that returns each
// UTF-8-encoded rune as a token. The sequence of runes returned is
// equivalent to that from a range loop over the input as a string, which
// means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd".
// Because of the Scan interface, this makes it impossible for the client to
// distinguish correctly encoded replacement runes from encoding errors.
func (r *Rune) Decode(data []byte, atEOF bool) (int, types.Object, error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }

    // Fast path 1: ASCII.
    if data[0] < utf8.RuneSelf {
        return 1, data[0:1], nil
    }

    // Fast path 2: Correct UTF-8 decode without error.
    _, width := utf8.DecodeRune(data)
    if width > 1 {
        // It's a valid encoding. Width cannot be one for a correctly encoded
        // non-ASCII rune.
        return width, data[0:width], nil
    }

    // We know it's an error: we have width==1 and implicitly r==utf8.RuneError.
    // Is the error because there wasn't a full rune to be decoded?
    // FullRune distinguishes correctly between erroneous and incomplete encodings.
    if !atEOF && !utf8.FullRune(data) {
        // Incomplete; get more bytes.
        return 0, nil, nil
    }

    // We have a real UTF-8 encoding error. Return a properly encoded error rune
    // but advance only one byte. This matches the behavior of a range loop over
    // an incorrectly encoded string.
    return 1, errorRune, nil
}

func init() {
    register.Codec(Namespace, New)
}
