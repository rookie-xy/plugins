package word_test

import (
	"testing"
	"unicode/utf8"
	"unicode"
	"strings"
)

var wordScanTests = []string{
	"",
	" ",
	"\n",
	"a",
	" a ",
	"abc def",
	" abc def ",
	" abc\tdef\nghi\rjkl\fmno\vpqr\u0085stu\u00a0\n",
}

// Test that the word splitter returns the same data as strings.Fields.
func TestScanWords(t *testing.T) {
	for n, test := range wordScanTests {
		buf := strings.NewReader(test)
		s := NewScanner(buf)
		s.Split(ScanWords)
		words := strings.Fields(test)
		var wordCount int
		for wordCount = 0; wordCount < len(words); wordCount++ {
			if !s.Scan() {
				break
			}
			got := s.Text()
			if got != words[wordCount] {
				t.Errorf("#%d: %d: expected %q got %q", n, wordCount, words[wordCount], got)
			}
		}
		if s.Scan() {
			t.Errorf("#%d: scan ran too long, got %q", n, s.Text())
		}
		if wordCount != len(words) {
			t.Errorf("#%d: termination expected at %d; got %d", n, len(words), wordCount)
		}
		err := s.Err()
		if err != nil {
			t.Errorf("#%d: %v", n, err)
		}
	}
}


// Test white space table matches the Unicode definition.
func TestSpace(t *testing.T) {
    for r := rune(0); r <= utf8.MaxRune; r++ {
        if IsSpace(r) != unicode.IsSpace(r) {
            t.Fatalf("white space property disagrees: %#U should be %t", r, unicode.IsSpace(r))
        }
    }
}



