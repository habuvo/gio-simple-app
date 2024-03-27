package main

import (
	"crypto/rand"
	"math/big"
	"time"
)

var (
	updatesTicker = make(chan ticker)
)

func update() {
	tc := time.NewTicker(time.Second * 2)

	for range tc.C {
		if len(myTickers) == 0 {
			continue
		}

		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(myTickers))))

		oldValue := myTickers[n.Int64()].value
		newValue := newValue(oldValue)
		diff := newValue - oldValue

		updatesTicker <- ticker{
			name:  myTickers[n.Int64()].name,
			value: newValue,
			diff:  diff,
		}
	}
}

func newValue(oldValue int) (newValue int) {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return oldValue
	}

	diff := ((float64(n.Int64()))/100 - 0.5) * float64(oldValue)

	return int(diff) + oldValue
}
