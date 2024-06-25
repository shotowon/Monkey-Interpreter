package object

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
	}

	return "NONE"
}
