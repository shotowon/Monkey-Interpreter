package object_test

import (
	"monkey/internal/object"
	"testing"
)

func TestHashMapKey(t *testing.T) {
	k1 := &object.String{Value: "key"}
	k2 := &object.String{Value: "key"}

	d1 := &object.String{Value: "diff"}
	d2 := &object.String{Value: "diff"}

	if k1.HashKey() != k2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if d1.HashKey() != d2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if k1.HashKey() == d1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
