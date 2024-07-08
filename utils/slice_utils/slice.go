package slice_utils

type IntInterface interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

type FloatInterface interface {
	float32 | float64
}

// Unique 切片去重
func Unique[T IntInterface | string](data []T) []T {
	list := make([]T, 0)
	hash := make(map[T]struct{})

	for _, value := range data {
		if _, ok := hash[value]; !ok {
			list = append(list, value)
			hash[value] = struct{}{}
		}
	}

	return list
}

func ToMap[T any, K int | string](arr []T, fn func(T) K) map[K]T {
	var m = make(map[K]T)

	for _, t := range arr {
		m[fn(t)] = t
	}

	return m
}
