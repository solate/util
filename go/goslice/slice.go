package goslice

// Intersect s1的交集
func Intersect[T comparable](s1, s2 []T) []T {
	inter := make([]T, 0)
	m := make(map[T]bool)

	// s1 存入map
	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}

	// s2 存在就添加到返回值
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			inter = append(inter, v)
		}
	}
	return inter
}

// Union 并集
func Union[T comparable](s1, s2 []T) []T {
	inter := make([]T, 0)
	m := make(map[T]bool)
	// s1 的值放入
	for _, v := range s1 {
		m[v] = true
		inter = append(inter, v)
	}
	// s2 不存在就添加
	for _, v := range s2 {
		if _, ok := m[v]; !ok {
			inter = append(inter, v)
		}
	}
	return inter
}

// Except s1的差集
func Except[T comparable](s1, s2 []T) []T {
	inter := make([]T, 0)
	m := make(map[T]bool)

	// s1 存入map
	for _, v := range s1 {
		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}
	// s2 删除相同的值
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			delete(m, v) // 存在就删除
		}
	}
	for key := range m {
		inter = append(inter, key)
	}
	return inter
}
