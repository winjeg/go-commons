package cmp

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func IfNumber[T Number](condition bool, a, b T) T {
	return If(condition, a, b) // 复用泛型函数
}

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// GT0OrLTEDefault greater than zero or less than default value
func GT0OrLTEDefault[T Number](n, def T) T {
	if n < 0 {
		return def
	}
	return min(n, def)
}
