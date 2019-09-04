package parser

type ym map[interface{}]interface{}

func asYm(raw interface{}) ym {
	m, ok := raw.(map[interface{}]interface{})
	if !ok {
		return nil
	}
	return ym(m)
}

func (m ym) readString(key string, fallbackValue string) string {
	v, ok := m[key].(string)
	if !ok {
		return fallbackValue
	}
	return v
}
