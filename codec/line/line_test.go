package line_test

import (
	"testing"
	"io"
	"bytes"
	"strings"
)

// slowReader is a reader that returns only a few bytes at a time, to test the incremental
// reads in Scanner.Scan.
type slowReader struct {
	max int
	buf io.Reader
}

func (sr *slowReader) Read(p []byte) (n int, err error) {
	if len(p) > sr.max {
		p = p[0:sr.max]
	}
	return sr.buf.Read(p)
}

// genLine writes to buf a predictable but non-trivial line of text of length
// n, including the terminal newline and an occasional carriage return.
// If addNewline is false, the \r and \n are not emitted.
func genLine(buf *bytes.Buffer, lineNum, n int, addNewline bool) {
	buf.Reset()
	doCR := lineNum%5 == 0
	if doCR {
		n--
	}
	for i := 0; i < n-1; i++ { // Stop early for \n.
		c := 'a' + byte(lineNum+i)
		if c == '\n' || c == '\r' { // Don't confuse us.
			c = 'N'
		}
		buf.WriteByte(c)
	}
	if addNewline {
		if doCR {
			buf.WriteByte('\r')
		}
		buf.WriteByte('\n')
	}
	return
}

// Test the line splitter, including some carriage returns but no long lines.
func TestScanLongLines(t *testing.T) {
	// Build a buffer of lots of line lengths up to but not exceeding smallMaxTokenSize.
	tmp := new(bytes.Buffer)
	buf := new(bytes.Buffer)
	lineNum := 0
	j := 0
	for i := 0; i < 2*smallMaxTokenSize; i++ {
		genLine(tmp, lineNum, j, true)
		if j < smallMaxTokenSize {
			j++
		} else {
			j--
		}
		buf.Write(tmp.Bytes())
		lineNum++
	}
	s := NewScanner(&slowReader{1, buf})
	s.Split(ScanLines)
	s.MaxTokenSize(smallMaxTokenSize)
	j = 0
	for lineNum := 0; s.Scan(); lineNum++ {
		genLine(tmp, lineNum, j, false)
		if j < smallMaxTokenSize {
			j++
		} else {
			j--
		}
		line := tmp.String() // We use the string-valued token here, for variety.
		if s.Text() != line {
			t.Errorf("%d: bad line: %d %d\n%.100q\n%.100q\n", lineNum, len(s.Bytes()), len(line), s.Text(), line)
		}
	}
	err := s.Err()
	if err != nil {
		t.Fatal(err)
	}
}

// Test that the line splitter errors out on a long line.
func TestScanLineTooLong(t *testing.T) {
	const smallMaxTokenSize = 256 // Much smaller for more efficient testing.
	// Build a buffer of lots of line lengths up to but not exceeding smallMaxTokenSize.
	tmp := new(bytes.Buffer)
	buf := new(bytes.Buffer)
	lineNum := 0
	j := 0
	for i := 0; i < 2*smallMaxTokenSize; i++ {
		genLine(tmp, lineNum, j, true)
		j++
		buf.Write(tmp.Bytes())
		lineNum++
	}
	s := NewScanner(&slowReader{3, buf})
	s.Split(ScanLines)
	s.MaxTokenSize(smallMaxTokenSize)
	j = 0
	for lineNum := 0; s.Scan(); lineNum++ {
		genLine(tmp, lineNum, j, false)
		if j < smallMaxTokenSize {
			j++
		} else {
			j--
		}
		line := tmp.Bytes()
		if !bytes.Equal(s.Bytes(), line) {
			t.Errorf("%d: bad line: %d %d\n%.100q\n%.100q\n", lineNum, len(s.Bytes()), len(line), s.Bytes(), line)
		}
	}
	err := s.Err()
	if err != ErrTooLong {
		t.Fatalf("expected ErrTooLong; got %s", err)
	}
}

// Test that the line splitter handles a final line without a newline.
func testNoNewline(text string, lines []string, t *testing.T) {
	buf := strings.NewReader(text)
	s := NewScanner(&slowReader{7, buf})
	s.Split(ScanLines)
	for lineNum := 0; s.Scan(); lineNum++ {
		line := lines[lineNum]
		if s.Text() != line {
			t.Errorf("%d: bad line: %d %d\n%.100q\n%.100q\n", lineNum, len(s.Bytes()), len(line), s.Bytes(), line)
		}
	}
	err := s.Err()
	if err != nil {
		t.Fatal(err)
	}
}

// Test that the line splitter handles a final line without a newline.
func TestScanLineNoNewline(t *testing.T) {
	const text = "abcdefghijklmn\nopqrstuvwxyz"
	lines := []string{
		"abcdefghijklmn",
		"opqrstuvwxyz",
	}
	testNoNewline(text, lines, t)
}

// Test that the line splitter handles a final line with a carriage return but no newline.
func TestScanLineReturnButNoNewline(t *testing.T) {
	const text = "abcdefghijklmn\nopqrstuvwxyz\r"
	lines := []string{
		"abcdefghijklmn",
		"opqrstuvwxyz",
	}
	testNoNewline(text, lines, t)
}

// Test that the line splitter handles a final empty line.
func TestScanLineEmptyFinalLine(t *testing.T) {
	const text = "abcdefghijklmn\nopqrstuvwxyz\n\n"
	lines := []string{
		"abcdefghijklmn",
		"opqrstuvwxyz",
		"",
	}
	testNoNewline(text, lines, t)
}

// Test that the line splitter handles a final empty line with a carriage return but no newline.
func TestScanLineEmptyFinalLineWithCR(t *testing.T) {
	const text = "abcdefghijklmn\nopqrstuvwxyz\n\r"
	lines := []string{
		"abcdefghijklmn",
		"opqrstuvwxyz",
		"",
	}
	testNoNewline(text, lines, t)
}
