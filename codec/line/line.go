package line

import (
    "bytes"
    "errors"

    "github.com/rookie-xy/hubble/types"
    "github.com/rookie-xy/hubble/log"
    "github.com/rookie-xy/hubble/codec"
    "github.com/rookie-xy/hubble/register"
  . "github.com/rookie-xy/hubble/log/level"
    "github.com/rookie-xy/hubble/adapter"
)

type Line struct {
    log.Log
    level  Level

    limit  int
    match  byte
}

func New(log log.Log, v types.Value) (codec.Decoder, error) {
    line := &Line{
        Log:   log,
        level: adapter.ToLevelLog(log).Get(),
        limit: -1,
        match: '\n',
    }

	if val := v.GetMap(); val != nil {
		if match, ok := val["match"]; ok {
            line.match = []byte(match.(string))[0]
        } else {
            return nil, errors.New("match is not found")
        }

        if max, ok := val["max"]; ok {
            line.limit = max.(int)
        } else {
            return nil, errors.New("max is not found")
        }
    }

    return line, nil
}

func (l *Line) Clone() types.Object {
    return &Line{
        Log:   l.Log,
        limit: l.limit,
        match: l.match,
    }
}

func (l *Line) Decode(in []byte) (types.Object, error) {
    return nil, nil
}

// Line is a split function for a Scanner that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
// The last non-empty line of configure will be returned even if it has no
// newline.
func (l *Line) LogDecode(data []byte, atEOF bool) (int, []byte, error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }

    if i := bytes.IndexByte(data, l.match); i >= 0 {
        // Out of bounds, throw out the line data
        if i > l.limit {
            l.log(WARN,"Out of bounds, throw out the line data")
            return i + 1, nil, nil
        }

        // We have a full newline-terminated line.
        return i + 1, dropCR(data[0:i]), nil
    }

    // If we're at EOF, we have a final, non-terminated line. Return it.
    if atEOF {
        return len(data), dropCR(data), nil
    }

    l.log(DEBUG, Name +"; request more data")
    return 0, nil, nil
}

func (l *Line) log(ll Level, fmt string, args ...interface{}) {
    log.Print(l.Log, l.level, ll, fmt, args...)
}

func init() {
    register.Decoder(Namespace, New)
}
