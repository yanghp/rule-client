package pkg

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"math/rand"
	"time"
)

func RandFloat(min, max float64, precision int) float64 {
	if min == max {
		return min
	}

	if max < min {
		min, max = max, min
	}
	rand.NewSource(time.Now().UnixMilli())
	n := rand.Float64()*(max-min) + min

	return mathutil.RoundToFloat(n, precision)
}
