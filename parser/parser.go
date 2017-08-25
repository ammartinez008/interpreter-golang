package parser

import (
	"github.com/monkey_lang/lexer"
	"github.com/monkey_lang/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}
