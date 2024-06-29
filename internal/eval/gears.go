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

func isErr(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.T_ERROR
	}

	return false
}

func applyFunc(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extEnv := extendFunctionEnv(fn, args)
		eval := Eval(fn.Body, extEnv)
		return unwrapReturnValue(eval)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}

}

func unwrapReturnValue(obj object.Object) object.Object {
	if ret, ok := obj.(*object.ReturnValue); ok {
		return ret.Value
	}

	return obj
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnv(fn.Env)

	for i := range fn.Params {
		env.Set(fn.Params[i].Value, args[i])
	}

	return env
}
