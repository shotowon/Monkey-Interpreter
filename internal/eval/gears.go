package eval

import "monkey/internal/object"

func boolToObj(value bool) *object.Boolean {
	if value {
		return TRUE
	}

	return FALSE
}

func isTrue(obj object.Object) bool {
	switch obj {
	case NULL, FALSE:
		return false
	case TRUE:
		return true
	default:
		return true
	}
}
