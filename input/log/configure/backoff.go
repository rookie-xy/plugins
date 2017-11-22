package configure

import "time"

type Backoff struct {
    Min     string
    Max     string

    min     time.Duration
    max     time.Duration
    Factor  int
}

func (b *Backoff) GetMax() time.Duration {
    return b.max
}

func (b *Backoff) GetMin() time.Duration {
    return b.min
}
