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
	"fmt"
)

// Log contains all log related data
type Log struct {
    conf        *Configure
	source       source.Source
	offset       int64
	lastTimeRead time.Time
    backoff      time.Duration
	log          log.Log
	done         chan struct{}
}

// New creates a new log instance to read log sources
func New(l log.Log, v types.Value) (input.Input, error) {
	log := &Log{
		lastTimeRead: time.Now(),
		log:          l,
		done:         make(chan struct{}),
	}

    if err := Init(v, log); err != nil {
    	return nil, err
	}

	return log, nil
}

func (f *Log) Init(src source.Source) error {
    var offset int64
	if seeker, ok := src.(io.Seeker); ok {
		var err error
		offset, err = seeker.Seek(0, os.SEEK_CUR)
		if err != nil {
			return err
		}
	}

	f.source = src
	f.offset = offset
	return nil
}

// Read reads from the reader and updates the offset
// The total number of bytes read is returned.
func (f *Log) Read(buf []byte) (int, error) {
	totalN := 0

	for {
		select {
		case <-f.done:
			return 0, source.ErrClosed
		default:
		}

		n, err := f.source.Read(buf)
		fmt.Println(string(buf))
		if n > 0 {
			f.offset += int64(n)
			f.lastTimeRead = time.Now()
		}
		totalN += n

		// Read from input completed without error
		// Either end reached or buffer full
		if err == nil {
			// reset backoff for next read
			f.backoff = f.conf.Min
			return totalN, nil
		}

		// Move buffer forward for next read
		buf = buf[n:]

		// Checks if an error happened or buffer is full
		// If buffer is full, cannot continue reading.
		// Can happen if n == bufferSize + io.EOF error
		err = f.errorChecks(err)
		if err != nil || len(buf) == 0 {
			return totalN, err
		}

		//logp.Debug("harvester", "End of file reached: %s; Backoff now.", f.fs.Name())
		f.wait()
	}
}

// errorChecks checks how the given error should be handled based on the config options
func (f *Log) errorChecks(err error) error {
	if err != io.EOF {
		//logp.Err("Unexpected state reading from %s; error: %s", f.fs.Name(), err)
		return err
	}

	// Stdin is not continuable
	if !f.source.Continuable() {
		//logp.Debug("harvester", "Source is not continuable: %s", f.fs.Name())
		return err
	}

	if err == io.EOF && f.conf.EOF {
		return err
	}

	// Refetch fileinfo to check if the file was truncated or disappeared.
	// Errors if the file was removed/rotated after reading and before
	// calling the stat function
	info, statErr := f.source.Stat()
	if statErr != nil {
		//logp.Err("Unexpected error reading from %s; error: %s", f.fs.Name(), statErr)
		return statErr
	}

	// check if file was truncated
	if info.Size() < f.offset {
		//logp.Debug("harvester",
		//	"File was truncated as offset (%d) > size (%d): %s", f.offset, info.Size(), f.fs.Name())
		return source.ErrFileTruncate
	}

	// Check file wasn't read for longer then CloseInactive
	age := time.Since(f.lastTimeRead)
	if age > f.conf.Inactive {
		return source.ErrInactive
	}

	if f.conf.Renamed {
		// Check if the file can still be found under the same path
		if !IsSameFile(f.source.Name(), info) {
			return source.ErrRenamed
		}
	}

	if f.conf.Removed {
		// Check if the file name exists. See https://github.com/elastic/filebeat/issues/93
		_, statErr := os.Stat(f.source.Name())

		// Error means file does not exist.
		if statErr != nil {
			return source.ErrRemoved
		}
	}

	return nil
}

func (f *Log) wait() {
	// Wait before trying to read file again. File reached EOF.
	select {
	case <-f.done:
		return
	case <-time.After(f.backoff):
	}

	// Increment backoff up to maxBackoff
	if f.backoff < f.conf.Max {
		f.backoff = f.backoff * time.Duration(f.conf.Factor)
		if f.backoff > f.conf.Max {
			f.backoff = f.conf.Max
		}
	}
}

// Close closes the done channel but no th the file handler
func (f *Log) Close() error {
	close(f.done)
	// Note: File reader is not closed here because that leads to race conditions
	return nil
}

func init() {
    register.Input(Namespace, New)
}
