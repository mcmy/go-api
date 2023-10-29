package utils

func CString(in any, def ...string) string {
	if out, ok := in.(string); ok {
		return out
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

func CInt(in any, def ...int) int {
	if out, ok := in.(int); ok {
		return out
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}
