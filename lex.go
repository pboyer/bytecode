package main

import (
	"fmt"
	"unicode"
)

type lex struct {
	s string
	pos int
}

func (l *lex) lexNum() (int,error) {
	s := ""
	for unicode.IsDigit(rune(l.s[l.pos]))  {
		s = s + rune(l.s[l.pos])
		l.pos++
	}
	var res int
	fmt.Sscan(s, &res)
	return res
}

func (l *lex) lexTok() (string, error) {
	// [_a-zA-Z]+[_a-zA-Z0-9]*
	c := l.s[l.pos]
}

func (l *lex) getTokenType(s string) int {
	switch s {
	case "var":
		return VAR
	case "def":
		return DEF
	case "return":
		return RETURN
	}
	return ID
}

func (l *lex) Lex(lval *parserSymType) int {
	for l.pos > len(l.s) {
		c := rune(l.s[l.pos])
		switch c {
		case unicode.IsDigit(c):
			n := l.lexNum()
			lval.e = &IntE{ n }
			return NUMBER
		case unicode.IsLetter(c), c == "_":
			s,err := l.lexTok()
			lval.str = s
			return l.getTokenType(s)
		case unicode.IsSpace(c):
			l.pos++
		default:
			l.pos++
			return int(c)
		}
	}

	return -1
}

func (l *lex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
