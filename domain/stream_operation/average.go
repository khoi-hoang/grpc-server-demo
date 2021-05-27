package stream_operation

import "errors"

type Average struct {
	sum   int64
	count int64
}

func (a *Average) Feed(num int64) {
	a.sum += num
	a.count++
}

func (a *Average) Fetch() (int64, error) {
	if a.count == 0 {
		return 0, errors.New("could not calculate average from nothing")
	}

	return a.sum / a.count, nil
}
