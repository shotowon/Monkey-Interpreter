package eval

import (
	"fmt"
	"monkey/internal/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type().String())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments to `first`. got=%d, expected=%d", len(args), 1)
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) > 0 {
					return arg.Elements[0]
				}

				return NULL
			default:
				return newError("unsupported argument passed to `first`. got=%s", args[0].Type().String())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments passed to `last`. got=%d, expected=%d", len(args), 1)
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len := len(arg.Elements); len > 0 {
					return arg.Elements[len-1]
				}

				return NULL
			default:
				return newError("unsupported argument passed to `last`. got=%s", args[0].Type().String())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments passed to `rest`. got=%d, expected=%d", len(args), 1)
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) > 0 {
					newArray := make([]object.Object, len(arg.Elements)-1, len(arg.Elements)-1)
					copy(newArray, arg.Elements[1:])
					return &object.Array{Elements: newArray}
				}

				return NULL
			default:
				return NULL
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, expected >= %d", len(args), 2)
			}

			switch arg := args[0].(type) {
			case *object.Array:
				arrLen := len(arg.Elements)
				argsLen := len(args) - 1
				newArrayLen := arrLen + argsLen
				newArray := make([]object.Object, newArrayLen, newArrayLen)
				copy(newArray, arg.Elements)
				for i := arrLen; i < newArrayLen; i++ {
					newArray[i] = args[i-arrLen+1]
				}

				return &object.Array{Elements: newArray}
			default:
				return NULL
			}
		},
	},
	"print": {
		Fn: func(args ...object.Object) object.Object {
			printable := ""
			for i := 0; i < len(args); i++ {
				printable += args[i].Inspect() + " "
			}
			fmt.Println(printable)

			return NULL
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for i := 0; i < len(args); i++ {
				fmt.Println(args[i].Inspect())
			}

			return NULL
		},
	},
}
