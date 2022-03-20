package common

import (
	"math/rand"
	"time"
)

func SetRandSeed() {
	rand.Seed(time.Now().UnixNano())
}
