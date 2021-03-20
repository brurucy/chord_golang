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

func ShouldContainValue(id int32, key int32, predId int32) bool {

	if key > predId && key <= id {
		return true
	} else if key < predId && predId > id && key <= id {
		return true
	} else if key > predId && predId >= id {
		return true
	} else if key > predId && id == predId {
		return true
	}

	return false
}

func ShouldContainValueTwo(id int32, key int32, predId int32) bool {

	if predId < key && key <= id {

		return true

	} else if predId < key && predId > id && key <= id {

		return true

	} else if predId < key && predId > id && key >= id {

		return true

	} else if predId > key && predId > id && key <= id {

		return true
	}
	return false
}
