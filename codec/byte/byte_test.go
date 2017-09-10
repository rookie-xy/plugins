package byte_test

import (
	"testing"
	"strings"
    "github.com/rookie-xy/hubble/scanner"
)

var scanTests = []string{
    "",
    "a",
    "¼",
    "☹",
    "\x81",   // UTF-8 error
    "\uFFFD", // correctly encoded RuneError
    "abcdefgh",
    "abc def\n\t\tgh    ",
    "abc¼☹\x81\uFFFD日本語\x82abc",
}

func TestScanByte(t *testing.T) {
    byte := byte.New(nil, nil)

    for n, test := range scanTests {
        buf := strings.NewReader(test)
        s := scanner.New(buf)
        s.Split(byte.Decode)

        var i int
        for i = 0; s.Scan(); i++ {
            if b := s.Bytes(); len(b) != 1 || b[0] != test[i] {
                t.Errorf("#%d: %d: expected %q got %q", n, i, test, b)
            }
        }

        if i != len(test) {
            t.Errorf("#%d: termination expected at %d; got %d", n, len(test), i)
        }

        err := s.Err()
        if err != nil {
            t.Errorf("#%d: %v", n, err)
        }
    }
}
