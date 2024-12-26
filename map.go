package utils

func Map[I any, L ~[]I, K comparable](items L, f func(I) K) map[K]I {
	m := make(map[K]I, len(items))
	for _, i := range items {
		m[f(i)] = i
	}
	return m
}

func MapE[I any, L ~[]I, K comparable](items L, f func(I) (K, error)) (map[K]I, error) {
	m := make(map[K]I, len(items))
	for _, i := range items {
		k, err := f(i)
		if err != nil {
			return nil, err
		}
		m[k] = i
	}
	return m, nil
}

func MapS[I any, L ~[]I, K comparable](items L, f func(I) K) []struct {
	K K
	I I
} {
	res := make([]struct {
		K K
		I I
	}, 0, len(items))
	for _, i := range items {
		res = append(res, struct {
			K K
			I I
		}{f(i), i})
	}
	return res
}

func MapSE[I any, L ~[]I, K comparable](items L, f func(I) (K, error)) ([]struct {
	K K
	I I
}, error) {
	res := make([]struct {
		K K
		I I
	}, 0, len(items))
	for _, i := range items {
		k, err := f(i)
		if err != nil {
			return nil, err
		}
		res = append(res, struct {
			K K
			I I
		}{k, i})
	}
	return res, nil
}

func GroupBy[I any, L ~[]I, K comparable](items L, f func(I) K) map[K]L {
	groups := make(map[K]L, len(items))
	for _, i := range items {
		key := f(i)
		g := groups[key]
		groups[key] = append(g, i)
	}
	return groups
}

func GroupByE[I any, L ~[]I, K comparable](items L, f func(I) (K, error)) (map[K]L, error) {
	groups := make(map[K]L, len(items))
	for _, i := range items {
		key, err := f(i)
		if err != nil {
			return nil, err
		}
		g := groups[key]
		groups[key] = append(g, i)
	}
	return groups, nil
}

func MapsIsEqual[K comparable, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valueA := range a {
		valueB, ok := b[key]
		if !ok || valueA != valueB {
			return false
		}
	}
	return true
}

func Intersection[I comparable, L ~[]I](items1 L, items2 L) map[I]struct{} {
	maxLen := len(items1)
	if len2 := len(items2); len2 > maxLen {
		maxLen = len2
	}

	items := make(map[I]struct{}, maxLen)
	for _, v1 := range items1 {
		for _, v2 := range items2 {
			if _, ok := items[v1]; !ok && v1 == v2 {
				items[v1] = struct{}{}
			}
		}
	}

	return items
}
