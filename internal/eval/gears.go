package eval

import "monkey/internal/object"

func boolToObj(value bool) *object.Boolean {
	if value {
		return TRUE
	}

	return FALSE
}
