package json

import (
	"bytes"
	"github.com/urso/go-structform/gotype"
    "github.com/urso/go-structform/json"
)

func reset(buffer bytes.Buffer, folder *gotype.Iterator) error {
	visitor := json.NewVisitor(&buffer)

	var err error
	folder, err = gotype.NewIterator(visitor,
		gotype.Folders(
			codec.MakeTimestampJson(),
			codec.MakeBCTimestampJson(),
		),
	)

	if err != nil {
		return err
	}

	return nil
}
