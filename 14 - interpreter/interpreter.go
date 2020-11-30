package interpreter

import (
	"errors"
	"strconv"
)

// Node 节点
type Node interface {
	Interpret() int
}

// Sign 符号接口
type Sign interface {
	Node
	// Priority 运算符优先级
	Priority() int
	// Set 设置二元运算符两侧数值
	Set(Node, Node)
}

// Symbol 左右括号等符号
type Symbol string

// Priority 优先级
func (s Symbol) Priority() int {
	switch s {
	case "(", ")":
		return 1
	}
	return 0
}

// Interpret 占位实现 Sign 接口
func (Symbol) Interpret() int { return 0 }

// Set 占位实现 Sign 接口
func (Symbol) Set(Node, Node) {}

var (
	// LeftParenthesis 左括号
	LeftParenthesis Symbol = "("
	// RightParenthesis 右括号
	RightParenthesis Symbol = ")"
)

// ValNode 数值节点
type ValNode struct {
	val int
}

// Interpret .
func (n *ValNode) Interpret() int {
	return n.val
}

// AddNode 加运算节点
type AddNode struct {
	left, right Node
}

// Set .
func (n *AddNode) Set(left, right Node) {
	n.left, n.right = left, right
}

// Priority .
func (n *AddNode) Priority() int {
	return 2
}

// Interpret .
func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

// SubNode 差运算节点
type SubNode struct {
	left, right Node
}

// Set .
func (n *SubNode) Set(left, right Node) {
	n.left, n.right = left, right
}

// Priority .
func (n *SubNode) Priority() int {
	return 2
}

// Interpret .
func (n *SubNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

// MulNode 乘法运算节点
type MulNode struct {
	left, right Node
}

// Set .
func (n *MulNode) Set(left, right Node) {
	n.left, n.right = left, right
}

// Priority .
func (n *MulNode) Priority() int {
	return 3
}

// Interpret .
func (n *MulNode) Interpret() int {
	return n.left.Interpret() * n.right.Interpret()
}

// DivNode 除法运算节点
type DivNode struct {
	left, right Node
}

// Set .
func (n *DivNode) Set(left, right Node) {
	n.left, n.right = left, right
}

// Priority .
func (n *DivNode) Priority() int {
	return 3
}

// Interpret .
func (n *DivNode) Interpret() int {
	return n.left.Interpret() / n.right.Interpret()
}

// Parser 解释器/分析器
type Parser struct {
	numStack    []Node
	symbolStack []Sign
	ret         Node
}

// Parse 解析
func (p *Parser) Parse(exp string) (err error) {
	for _, c := range exp {
		switch c {
		case ' ':
			continue
		case '+':
			p.symbolStack = append(p.symbolStack, p.NewAddNode())
		case '-':
			p.symbolStack = append(p.symbolStack, p.NewSubNode())
		case '*', '×', '✖':
			p.symbolStack = append(p.symbolStack, p.NewMulNode())
		case '/', '➗', '÷':
			p.symbolStack = append(p.symbolStack, p.NewDivNode())
		case '(':
			p.symbolStack = append(p.symbolStack, LeftParenthesis)
		case ')':
			for l := len(p.symbolStack); l > 0; l = len(p.symbolStack) {
				sign := p.symbolStack[l-1]
				if s, ok := sign.(Symbol); ok && s == LeftParenthesis {
					p.symbolStack = p.symbolStack[:l-1]
					break
				}

				p.PopSign()
			}
		default:
			n, err := strconv.Atoi(string(c))
			if err != nil {
				return err
			}
			p.numStack = append(p.numStack, &ValNode{val: n})
		}
	}
	return
}

// PopSign pop 符号栈
func (p *Parser) PopSign() {
	l, r := p.numStack[len(p.numStack)-2], p.numStack[len(p.numStack)-1]
	p.numStack = p.numStack[:len(p.numStack)-2]
	sign := p.symbolStack[len(p.symbolStack)-1]
	p.symbolStack = p.symbolStack[:len(p.symbolStack)-1]
	sign.Set(l, r)
	p.numStack = append(p.numStack, sign)
}

// NewAddNode 创建加法节点
func (p *Parser) NewAddNode() *AddNode {
	newNode := &AddNode{}
	if l := len(p.symbolStack); l > 0 && p.symbolStack[l-1].Priority() >= newNode.Priority() {
		p.PopSign()
	}
	return newNode
}

// NewSubNode 创建减法节点
func (p *Parser) NewSubNode() *SubNode {
	newNode := &SubNode{}
	if l := len(p.symbolStack); l > 0 && p.symbolStack[l-1].Priority() >= newNode.Priority() {
		p.PopSign()
	}
	return newNode
}

// NewMulNode 创建乘法节点
func (p *Parser) NewMulNode() *MulNode {
	newNode := &MulNode{}
	if l := len(p.symbolStack); l > 0 && p.symbolStack[l-1].Priority() >= newNode.Priority() {
		p.PopSign()
	}

	return newNode
}

// NewDivNode 创建除法节点
func (p *Parser) NewDivNode() *DivNode {
	newNode := &DivNode{}
	if l := len(p.symbolStack); l > 0 && p.symbolStack[l-1].Priority() >= newNode.Priority() {
		p.PopSign()
	}
	return newNode
}

// Result 获取计算结果
func (p *Parser) Result() (n Node, err error) {
	for l := len(p.symbolStack); l > 0; l = len(p.symbolStack) {
		p.PopSign()
	}

	if len(p.numStack) != 1 {
		err = errors.New("err")
		return
	}
	n = p.numStack[0]
	return
}
