package mapper

// Map is a Generic JSON handler
type Map map[string]interface{}

// M returns a nested map
func (m Map) M(s string) Map {
	return m[s].(map[string]interface{})
}

// A returns an array
func (m Map) A(s string) []interface{} {
	return m[s].([]interface{})
}

// S returns a string
func (m Map) S(s string) string {
	return m[s].(string)
}

// B returns a bool
func (m Map) B(s string) bool {
	return m[s].(bool)
}
