package csvparse

import (
	"encoding/csv"
	"github.com/pkg/errors"
	"io"
)

// Unmarshaller is the interface implemented by types
// that can unmarshal a CSV description of themselves.
// The input can be assumed to be a valid encoding of a CSV value.
type Unmarshaller interface {
	UnmarshalCSV(in []string) error
}

// Marshaller is the interface implemented by types
// that can marshal themselves into valid CSV.
type Marshaller interface {
	MarshalCSV() ([]string, error)
}

// Alloc defines a function that allocates Unmarshaller memory.
// All pointer structer fields must be allocated internally.
type Alloc func() Unmarshaller

type result struct {
	n int64
	l []string
}

// ProcessCSV function to parse csv file.
func ProcessCSV(rc *csv.Reader) (<-chan []string, <-chan error) {
	resc := make(chan []string)
	errc := make(chan error, 1)
	go func() {
		for {
			rec, err := rc.Read()
			if err != nil {
				errc <- err
				return
			}
			if rec == nil {
				errc <- errors.New("Null read record")
				return
			}
			resc <- rec
		}
	}()
	return resc, errc
}

// ReadAll executes UnmarshalCSV for every line of csv file
func ReadAll(rc *csv.Reader, allocr Alloc) ([]Unmarshaller, error) {
	resc, errc := ProcessCSV(rc)

	rlist := []Unmarshaller{}
	for {
		select {
		case in := <-resc:
			// if in == nil {
			// 	continue
			// }
			row := allocr()
			if err := row.UnmarshalCSV(in); err != nil {
				return nil, errors.Wrap(err, "Error in UnmarshalCSV")
			}
			rlist = append(rlist, row)
		case err := <-errc:
			if err == io.EOF {
				return rlist, nil
			}
			return nil, errors.Wrap(err, "ReadAll csvparse error")
		}
	}
}
