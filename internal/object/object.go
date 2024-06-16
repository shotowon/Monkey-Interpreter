package object

type ObjectType uint

type Object interface {
	Type() ObjectType
	Inspect() string
}
