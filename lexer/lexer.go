package lexer

type Lexer struct {
	input        string
	position     int  //current position in input
	readPosition int  //cur position in input
	ch           byte //current char
}

//todo: support unicode and Emoji
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.input[l.readPosition]
	l.readPosition += 1
}
