package log

import (
	"time"
	"io"
	"os"
	"github.com/rookie-xy/hubble/log"
	"github.com/rookie-xy/hubble/types"
	"github.com/rookie-xy/hubble/source"
	"github.com/rookie-xy/hubble/register"
   "github.com/rookie-xy/hubble/input"
    "github.com/rookie-xy/plugins/input/log/configure"
	"github.com/rookie-xy/plugins/input/log/utils"
  . "github.com/rookie-xy/hubble/log/level"
	"github.com/rookie-xy/hubble/adapter"
)

// Log contains all log related data
type Log struct {
    configure     *configure.Configure
	source         source.Source
	offset         int64

	lastTimeRead   time.Time
    backoff        time.Duration
	log.Log
	level          Level
	done           chan struct{}
}

// New creates a new log instance to read log sources
func New(l log.Log, v types.Value) (input.Input, error) {
	log := &Log{
		lastTimeRead: time.Now(),
		Log:          l,
	    level:        adapter.ToLevelLog(l).Get(),
		done:         make(chan struct{}),
	}

	configure := configure.New()
    if err := configure.Init(v); err != nil {
    	return nil, err
	}

	log.configure = configure
	log.backoff = configure.Backoff.GetMin()
	return log, nil
}

func (l *Log) Clone() types.Object {
    return &Log{
    	configure: l.configure,
    	lastTimeRead: time.Now(),
    	backoff: l.backoff,
        Log: l.Log,
        level: l.level,
    	done: make(chan struct{}),
	}
}

func (l *Log) Init(src source.Source) error {
    var offset int64
	if seeker, ok := src.(io.Seeker); ok {
		var err error
		offset, err = seeker.Seek(0, os.SEEK_CUR)
		if err != nil {
			return err
		}
	}

	l.source = src
	l.offset = offset

	l.log(DEBUG, "input log init success")
	return nil
}

// Read reads from the reader and updates the offset
// The total number of bytes read is returned.
func (l *Log) Read(buf []byte) (int, error) {
	totalN := 0

	for {
		select {
		case <-l.done:
			return 0, source.ErrClosed
		default:
		}

		n, err := l.source.Read(buf)
		if n > 0 {
			l.offset += int64(n)
			l.lastTimeRead = time.Now()
		}
		totalN += n

		// Read from configure completed without error
		// Either end reached or buffer full
		if err == nil {
			// reset backoff for next read
			l.backoff = l.configure.Backoff.GetMin()
			return totalN, nil
		}

		// Move buffer forward for next read
		buf = buf[n:]

		// Checks if an error happened or buffer is full
		// If buffer is full, cannot continue reading.
		// Can happen if n == bufferSize + io.EOF error
		err = l.errorChecks(err)
		if err != nil || len(buf) == 0 {
			return totalN, err
		}

		l.log(DEBUG,"Collector End of file reached: %s; backoff now", l.source.Name())
		l.wait()
	}
}

// errorChecks checks how the given error should be handled based on the config options
func (l *Log) errorChecks(err error) error {
	if err != io.EOF {
		l.log(ERROR,"Unexpected models reading from %s; error: %s", l.source.Name(), err)
		return err
	}

	// Stdin is not continuable
	if !l.source.Continuable() {
        l.log(DEBUG,"harvester", "Source is not continuable: %s", l.source.Name())
		return err
	}

	if err == io.EOF && l.configure.EOF {
		return err
	}

	// Refetch fileinfo to check if the file was truncated or disappeared.
	// Errors if the file was removed/rotated after reading and before
	// calling the stat function
	info, statErr := l.source.Stat()
	if statErr != nil {
		l.log(ERROR,"Unexpected error reading from %s; error: %s", l.source.Name(), statErr)
		return statErr
	}

	// check if file was truncated
	if info.Size() < l.offset {
		l.log(DEBUG,"collector file was truncated as offset (%d) > size (%d): %s",
			             l.offset, info.Size(), l.source.Name())
		return source.ErrFileTruncate
	}

	// Check file wasn't read for longer then CloseInactive
	age := time.Since(l.lastTimeRead)
	if age > l.configure.GetInactive() {
		return source.ErrInactive
	}

	if l.configure.Renamed {
		// Check if the file can still be found under the same path
		if !utils.SameFile(l.source.Name(), info, l.log) {
			return source.ErrRenamed
		}
	}

	if l.configure.Removed {
		// Check if the file name exists. See https://github.com/elastic/filebeat/issues/93
		_, statErr := os.Stat(l.source.Name())

		// Error means file does not exist.
		if statErr != nil {
			return source.ErrRemoved
		}
	}

	return nil
}

func (l *Log) wait() {
	// Wait before trying to read file again. File reached EOF.
	select {
	case <-l.done:
		return
	case <-time.After(l.backoff):
	}

	// Increment backoff up to maxBackoff
	if backoff := l.configure.Backoff; l.backoff < backoff.GetMax() {
		l.backoff = l.backoff * time.Duration(backoff.Factor)
		if l.backoff > backoff.GetMax() {
			l.backoff = backoff.GetMax()
		}
	}
}

// Close closes the done channel but no th the file handler
func (l *Log) Close() error {
	close(l.done)
	// Note: File reader is not closed here because that leads to race conditions
	return nil
}

func (l *Log) log(ll Level, fmt string, args ...interface{}) {
    log.Print(l.Log, l.level, ll, fmt, args...)
}

func init() {
    register.Input(Namespace, New)
}
