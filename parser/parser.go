package parser

import (
	"goInterpreter/ast"
	"goInterpreter/lexer"
	"goInterpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token //現在のトークン
	peekToken token.Token //次のトークン
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {

}
