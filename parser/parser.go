package parser

import (
	"github.com/monkey_lang/ast"
	"github.com/monkey_lang/lexer"
	"github.com/monkey_lang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//Read first token
	// set peek token to next token
	curToken = l.NextToken
	peekToken = l.NextToken
	return p
}

func (p *Parser) nextToken() {
	curToken = peekToken
	peekToken = p.l.NextToken
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil

}
