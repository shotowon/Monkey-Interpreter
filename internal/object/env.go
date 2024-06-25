package object

type Environment struct {
	store map[string]Object
}

func NewEnv() *Environment {
	return &Environment{make(map[string]Object)}
}

func (e *Environment) Get(name string) (Object, bool) {
	val, ok := e.store[name]
	return val, ok
}

func (e *Environment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}
