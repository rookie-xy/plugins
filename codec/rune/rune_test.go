package rune_test

import (
	"testing"
	"strings"
	"unicode/utf8"
	"github.com/rookie-xy/hubble/scanner"
)

// Test that the rune splitter returns same sequence of runes (not bytes) as for range string.
func TestScanRune(t *testing.T) {
	rune := New(nil, nil)

	for n, test := range scanTests {
		buf := strings.NewReader(test)
		s := scanner.New(buf)
		s.Split(rune.Decode)

		var i, runeCount int
		var expect rune
		// Use a string range loop to validate the sequence of runes.
		for i, expect = range string(test) {
			if !s.Scan() {
				break
			}
			runeCount++
			got, _ := utf8.DecodeRune(s.Bytes())
			if got != expect {
				t.Errorf("#%d: %d: expected %q got %q", n, i, expect, got)
			}
		}

		if s.Scan() {
			t.Errorf("#%d: scan ran too long, got %q", n, s.Text())
		}

		testRuneCount := utf8.RuneCountInString(test)
		if runeCount != testRuneCount {
			t.Errorf("#%d: termination expected at %d; got %d", n, testRuneCount, runeCount)
		}

		err := s.Err()
		if err != nil {
			t.Errorf("#%d: %v", n, err)
		}
	}
}
