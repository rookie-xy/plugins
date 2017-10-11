package log

import "time"

type Configure struct {
   *Backoff

    Inactive  time.Duration
    Timeout   time.Duration

    Removed   bool
    Renamed   bool
    EOF       bool
}
