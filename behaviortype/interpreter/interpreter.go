/*
概念与定义：解释器模式（Interpreter Pattern）是一种行为型设计模式，用于定义一个语言的文法，并通过构建一个解释器来解释该语言中的句子。通过定义
一个抽象的表达式接口，为语言中的句子或表达式提供解释操作
应用场景：
	1.设计解析器或执行器，解释领域特定的语言或表达式。
	2.实现简单的语法解析器，如数学表达式、正则表达式解析、简单脚本语言等。
*/

package interpreter

import (
	"strconv"
	"strings"
)

// Expression 是抽象表达式接口，定义解释方法
type Expression interface {
	Interpret() int
}

// Number 是终端表达式，表示具体的数字
type Number struct {
	value int
}

// NewNumber 创建新的 Number 实例
func NewNumber(value int) *Number {
	return &Number{value: value}
}

func (n *Number) Interpret() int {
	return n.value
}

// AddExpression 表示加法操作的非终端表达式
type AddExpression struct {
	left, right Expression
}

// NewAddExpression 创建新的加法表达式
func NewAddExpression(left, right Expression) *AddExpression {
	return &AddExpression{left: left, right: right}
}

func (a *AddExpression) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

// SubtractExpression 表示减法操作的非终端表达式
type SubtractExpression struct {
	left, right Expression
}

// NewSubtractExpression 创建新的减法表达式
func NewSubtractExpression(left, right Expression) *SubtractExpression {
	return &SubtractExpression{left: left, right: right}
}

func (s *SubtractExpression) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}

// Parser 是解析器，用于解析和构建表达式树
type Parser struct {
	expression Expression
}

// NewParser 创建新的 Parser 实例
func NewParser(input string) *Parser {
	tokens := strings.Fields(input)
	stack := []Expression{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token {
		case "+":
			left := stack[len(stack)-2]
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, NewAddExpression(left, right))
		case "-":
			left := stack[len(stack)-2]
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, NewSubtractExpression(left, right))
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, NewNumber(num))
		}
	}
	return &Parser{expression: stack[0]}
}

// Interpret 解析表达式
func (p *Parser) Interpret() int {
	return p.expression.Interpret()
}
