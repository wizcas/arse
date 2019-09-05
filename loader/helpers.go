package loader

type freemap map[interface{}]interface{}

func asYm(raw interface{}) freemap {
	m, ok := raw.(map[interface{}]interface{})
	if !ok {
		return nil
	}
	return freemap(m)
}

func (m freemap) readString(key string, fallbackValue string) string {
	v, ok := m[key].(string)
	if !ok {
		return fallbackValue
	}
	return v
}
