package eval

import (
	"fmt"
	"monkey/internal/ast"
	"monkey/internal/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.BooleanExpression:
		return boolToObj(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isErr(right) {
			return right
		}

		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isErr(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isErr(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BlockStatement:
		return evalBlockStmt(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isErr(val) {
			return val
		}

		env.Set(node.Name.Value, val)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isErr(val) {
			return val
		}

		return &object.ReturnValue{Value: val}
	}

	return nil
}

func evalInfixExpression(op string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.T_INTEGER && right.Type() == object.T_INTEGER:
		return evalIntegerInfixExpression(op, left, right)
	case op == "==":
		return boolToObj(left == right)
	case op == "!=":
		return boolToObj(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), op, right.Type())
	}

	return newError("unknown operator: %s %s %s", left.Type(), op, right.Type())
}

func evalIntegerInfixExpression(op string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch op {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return boolToObj(leftVal < rightVal)
	case ">":
		return boolToObj(leftVal > rightVal)
	case "<=":
		return boolToObj(leftVal <= rightVal)
	case ">=":
		return boolToObj(leftVal >= rightVal)
	case "==":
		return boolToObj(leftVal == rightVal)
	case "!=":
		return boolToObj(leftVal != rightVal)
	}

	return newError("unknown operator: %s %s %s", left.Type(), op, right.Type())
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOpExpression(right)
	case "-":
		return evalMinusPrefixOpExpression(right)
	}

	return newError("uknown operator: %s%s", operator, right.Type())
}

func evalMinusPrefixOpExpression(right object.Object) object.Object {
	if right.Type() != object.T_INTEGER {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalBangOpExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	}

	return FALSE
}

func evalStatements(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt, env)

		if ret, ok := result.(*object.ReturnValue); ok {
			return ret.Value
		}
	}

	return result
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	cond := Eval(ie.Condition, env)
	if isErr(cond) {
		return cond
	}

	if isTrue(cond) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	}

	return NULL
}

func evalProgram(p *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range p.Statements {
		result = Eval(stmt, env)

		switch result := result.(type) {
		case *object.Error:
			return result
		case *object.ReturnValue:
			return result.Value
		}
	}

	return result
}

func evalBlockStmt(b *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range b.Statements {
		result = Eval(stmt, env)

		if result != nil {

			if rt := result.Type(); rt == object.T_RETURN_VALUE || rt == object.T_ERROR {
				return result
			}
		}
	}

	return result
}

func newError(format string, args ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, args...)}
}
