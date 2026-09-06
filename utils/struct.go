package utils

// AnyGet 从 map[string]any（或指向它的指针、指向它的 *any）中读取字段 f 并断言为 K；任何一步失败都返回零值。
func AnyGet[K any](d any, f string) K {
	var zero K
	m, ok := asMap(d)
	if !ok {
		return zero
	}
	v, ok := m[f].(K)
	if !ok {
		return zero
	}
	return v
}

// AnySet 向 map[string]any 写入字段。t 必须是指向该 map 的指针（*map[string]any，或指向 map 的 *any），否则返回 false。
func AnySet(t, d any, fieldName string) bool {
	var m map[string]any
	switch v := t.(type) {
	case *map[string]any:
		if v != nil {
			m = *v
		}
	case *any:
		if v != nil {
			m, _ = asMap(*v)
		}
	}
	if m == nil {
		return false
	}
	m[fieldName] = d
	return true
}

func asMap(v any) (map[string]any, bool) {
	switch x := v.(type) {
	case map[string]any:
		return x, true
	case *map[string]any:
		if x != nil {
			return *x, true
		}
	case *any:
		if x != nil {
			return asMap(*x)
		}
	}
	return nil, false
}
