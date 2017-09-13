package multiline

type line struct {
    data []byte
}

func (c *line) Concat() *line {
    return nil
}

func (c *line) Length() int {
    return nil
}

func (c *line) Clear() int {
    return nil
}

func (c *line) Get() []byte {
    return nil
}
