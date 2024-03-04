package randomNum

import (
	"math/rand"
)

func GetRandomNumInComposition(max int, min int) (int, error) {
	if max < min {
		return 0, ErrorMaxValue{}
	}

	return rand.Intn(max-min) + min, nil
}
