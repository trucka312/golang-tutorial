

package sort

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/max"
	"github.com/TheAlgorithms/Go/math/min"
)


func Pigeonhole[T constraints.Integer](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}

	max := max.Int(arr...)
	min := min.Int(arr...)

	size := max - min + 1

	holes := make([]T, size)

	for _, element := range arr {
		holes[element-min]++
	}

	i := 0

	for j := T(0); j < size; j++ {
		for holes[j] > 0 {
			holes[j]--
			arr[i] = j + min
			i++
		}
	}

	return arr
}
