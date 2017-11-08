package dump

import "os"

const (
    flag  = os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_SYNC
    model = 0600
)
