package eval

import (
	"monkey/internal/ast"
	"monkey/internal/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.BooleanExpression:
		return boolToObj(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.BlockStatement:
		return evalBlockStmt(node)
	case *ast.IfExpression:
		return evalIfExpression(node)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
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
	}

	return NULL
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

	return NULL
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOpExpression(right)
	case "-":
		return evalMinusPrefixOpExpression(right)
	}

	return NULL
}

func evalMinusPrefixOpExpression(right object.Object) object.Object {
	if right.Type() != object.T_INTEGER {
		return NULL
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

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)

		if ret, ok := result.(*object.ReturnValue); ok {
			return ret.Value
		}
	}

	return result
}

func evalIfExpression(ie *ast.IfExpression) object.Object {
	cond := Eval(ie.Condition)
	if isTrue(cond) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	}

	return NULL
}

func evalProgram(p *ast.Program) object.Object {
	var result object.Object

	for _, stmt := range p.Statements {
		result = Eval(stmt)

		if ret, ok := result.(*object.ReturnValue); ok {
			return ret.Value
		}
	}

	return result
}

func evalBlockStmt(b *ast.BlockStatement) object.Object {
	var result object.Object

	for _, stmt := range b.Statements {
		result = Eval(stmt)

		if result != nil && result.Type() == object.T_RETURN_VALUE {
			return result
		}
	}

	return result
}
