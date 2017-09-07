package byte

// ScanBytes is a split function for a Scanner that returns each byte as a token.
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }
    return 1, data[0:1], nil
}
