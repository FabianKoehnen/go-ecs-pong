package util

import "golang.org/x/exp/constraints"

func Between[T constraints.Ordered](num, min, max T) bool {
	return num >= min && num <= max
}
