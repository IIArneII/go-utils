package utils

func P[T any](v T) *T {
	return &v
}

func T[T any](value bool, trueValue T, falseValue T) T {
	if value {
		return trueValue
	}
	return falseValue
}

func TF[T any](value bool, trueValue func() T, falseValue func() T) T {
	if value {
		return trueValue()
	}
	return falseValue()
}

func Contains[I comparable, L ~[]I](item I, items L) bool {
	for _, i := range items {
		if item == i {
			return true
		}
	}
	return false
}

func Select[I any, L ~[]I, T any](items L, f func(I) T) []T {
	res := make([]T, 0, len(items))
	for _, i := range items {
		res = append(res, f(i))
	}
	return res
}

func SelectE[I any, L ~[]I, T any](items L, f func(I) (T, error)) ([]T, error) {
	res := make([]T, 0, len(items))
	for _, i := range items {
		r, err := f(i)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func Find[I any, L ~[]I](items L, f func(I) bool) (i I, ok bool) {
	for _, v := range items {
		if f(v) {
			return v, true
		}
	}

	return i, false
}

func FindE[I any, L ~[]I](items L, f func(I) (bool, error)) (i I, ok bool, err error) {
	for _, v := range items {
		ok, err := f(v)
		if err != nil {
			return i, false, err
		}
		if ok {
			return v, true, nil
		}
	}
	return i, false, nil
}
