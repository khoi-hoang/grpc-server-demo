package stream_operation

import "io"

type primeDecomposer struct{
	num   int64
	prime int64
}

func NewPrimeDecomposer() *primeDecomposer {
	return &primeDecomposer{prime: 2}
}

func (p *primeDecomposer) Feed(num int64) {
	p.num = num
}

func (p *primeDecomposer) Fetch() (int64, error) {
	for p.num > 1 {
		if p.num % p.prime == 0 {
			p.num /= p.prime
			return p.prime, nil
		} else {
			p.prime++
		}
	}

	return 0, io.EOF
}

