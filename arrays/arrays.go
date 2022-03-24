package arrays

type number interface {
	int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		int | uint | float32 | float64
}

func Sum[X number](arr []X) X {
	sum := X(0)
	for _, n := range arr {
		sum += n
	}
	return sum
}

func Contains[X number | bool | string](arr []X, found X) bool {
	for _, val := range arr {
		if found == val {
			return true
		}
	}
	return false
}

func Remove[X number | bool | string](arr []X, deleted X) []X {
	index := FindIndex(arr, deleted)
	if index == -1 {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

func Equals[X number | bool | string](arr1, arr2 []X) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func EqualsItems[X number | bool | string](arr1, arr2 []X) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		index := FindIndex(arr2, arr1[i])
		if index == -1 {
			return false
		}
	}
	return true
}

func FindIndex[X number | bool | string](arr []X, search X) int {
	for index, val := range arr {
		if val == search {
			return index
		}
	}
	return -1
}
