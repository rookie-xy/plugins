package log

import "time"

type Backoff struct {
    Min     time.Duration
    Max     time.Duration
    Factor  int
}
/*
func (b *Backoff) Check(current *time.Duration) {
 	// Increment backoff up to maxBackoff
 	backoff := *current
	if backoff < b.Max {
		current = backoff * time.Duration(b.Factor)
		if backoff > b.Max {
			backoff = b.Max
		}
	}
}
*/
