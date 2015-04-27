package ast

type Expression interface {
	CanBeExpression()
	// String() string
}

type UnaryExpression struct {
	Op   *Op
	Expr Expression
}

func NewUnaryExpression(op *Op, expr Expression) *UnaryExpression {
	return &UnaryExpression{op, expr}
}

func (self UnaryExpression) CanBeExpression() {}

type BinaryExpression struct {
	Op    *Op
	Left  Expression
	Right Expression
}

func (self BinaryExpression) CanBeExpression() {}

func NewBinaryExpression(op *Op, left Expression, right Expression) *BinaryExpression {
	return &BinaryExpression{op, left, right}
}