package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnv() *Environment {
	return &Environment{make(map[string]Object), nil}
}

func NewEnclosedEnv(outer *Environment) *Environment {
	env := NewEnv()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	val, ok := e.store[name]
	if !ok && e.outer != nil {
		val, ok = e.outer.Get(name)
	}
	return val, ok
}

func (e *Environment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}
