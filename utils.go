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

func Filter[I any, L ~[]I](items L, f func(I) bool) []I {
	res := make([]I, 0, len(items)/2)
	for _, i := range items {
		if f(i) {
			res = append(res, i)
		}
	}
	return res
}

func FilterE[I any, L ~[]I](items L, f func(I) (bool, error)) ([]I, error) {
	res := make([]I, 0, len(items)/2)
	for _, i := range items {
		ok, err := f(i)
		if err != nil {
			return nil, err
		}
		if ok {
			res = append(res, i)
		}
	}
	return res, nil
}

func All[I any, L ~[]I](items L, f func(I) bool) bool {
	for _, i := range items {
		if !f(i) {
			return false
		}
	}
	return true
}

func AllE[I any, L ~[]I](items L, f func(I) (bool, error)) (bool, error) {
	for _, i := range items {
		ok, err := f(i)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

func Any[I any, L ~[]I](items L, f func(I) bool) bool {
	for _, i := range items {
		if f(i) {
			return true
		}
	}
	return false
}

func AnyE[I any, L ~[]I](items L, f func(I) (bool, error)) (bool, error) {
	for _, i := range items {
		ok, err := f(i)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}

func Reduce[I any, L ~[]I, R any](items L, f func(R, I) R, initial R) R {
	res := initial
	for _, i := range items {
		res = f(res, i)
	}
	return res
}

func ReduceE[I any, L ~[]I, R any](items L, f func(R, I) (R, error), initial R) (R, error) {
	var err error

	res := initial
	for _, i := range items {
		res, err = f(res, i)
		if err != nil {
			return initial, err
		}
	}
	return res, nil
}
