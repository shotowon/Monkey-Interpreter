package object

func (o ObjectType) String() string {
	switch o {
	case T_INTEGER:
		return "INTEGER"
	}

	return "NONE"
}
