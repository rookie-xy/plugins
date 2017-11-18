package input

import "time"

type Backoff struct {
    Min     time.Duration
    Max     time.Duration
    Factor  int
}
