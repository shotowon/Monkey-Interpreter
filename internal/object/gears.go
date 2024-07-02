package object

import "hash/fnv"

func (o ObjectType) String() string {
	switch o {
	case T_INTEGER:
		return "INTEGER"
	case T_BOOL:
		return "BOOL"
	case T_NULL:
		return "NULL"
	case T_RETURN_VALUE:
		return "RETURN VALUE"
	case T_ERROR:
		return "ERROR"
	case T_STRING:
		return "STRING"
	case T_BUILTIN:
		return "BUILTIN"
	}

	return "NONE"
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (b *Boolean) HashKey() HashKey {
	var val uint64

	if b.Value {
		val = 1
	} else {
		val = 0
	}

	return HashKey{Type: b.Type(), Value: val}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
