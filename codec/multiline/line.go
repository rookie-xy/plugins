package multiline

type line struct {
    data []byte
}

func (c *line) Concat(b []byte) *line {
    return nil
}

func (c *line) Length() int {
    return -1
}

func (c *line) Clear() int {
    return -1
}

func (c *line) Get() []byte {
    return nil
}
