package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	//Special
	ILLEGAL = "ILLEGAL" // illegal character
	EOF     = "EOF"     // end of file

	//identifiers and literals
	IDENT = "IDENT" //identifiers, function names
	INT   = "INT"   // int values

	//Operators
	ASSIGN = "=" // ASSIGN value
	PLUS   = "+" // addition

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}
