package ast

import "goInterpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier //束縛の識別子を保持
	Value Expression  //値を生成する式を保持する為
}

type ReturnStatement struct {
	Token token.Token		//returnトークン
	ReturnValue Expression
}

//Statement interfaceを満たす
func (ls *LetStatement) statementNode() {}

//Node interfaceを満たす
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

//let x = 5;　におけるxのような識別子を保持する為
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//Return
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {return rs.Token.Literal}