package eval

import (
	"monkey/internal/ast"
	"monkey/internal/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
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
	}

	return nil
}

func evalInfixExpression(op string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.T_INTEGER && right.Type() == object.T_INTEGER:
		return evalIntegerInfixExpression(op, left, right)
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
	}

	return result
}
