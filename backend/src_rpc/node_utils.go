package src_rpc


func AbsInt(x int32) int32 {

	if x < 0 {

		return -x

	}
	return x

}

func Mod(a, b int32) int32 {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func Min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func RingDistance(from, to, maxSize, minSize int32) int32 {

	toFrom := to - from
	maxSizeToFrom := AbsInt((maxSize - (to - from)))

	result := Min(toFrom, maxSizeToFrom)

	if to > from {

		return result

	} else if from == to {

		return 0

	} else {

		result = Mod((maxSize - minSize + result), maxSize)
		return result

	}

}

