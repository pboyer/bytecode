package main

import (
	"fmt"
	"unicode"
)

type lex struct {
	s string
	pos int
	result *BlockS
}

func (l *lex) lexNum() (int,error) {
	// [0-9]+

	s := ""
	for l.pos < len(l.s) {
		c := rune(l.s[l.pos])
		switch {
		case unicode.IsDigit(c):
			s += string(c)
			l.pos++
		default:
			goto done
		}
	}
	done:

	var res int
	_, err := fmt.Sscan(s, &res)
	if err != nil {
		return -1, err
	}
	return res, nil
}

func (l *lex) lexTok() (string, error) {
	// [a-zA-Z]+[a-zA-Z0-9]*

	s := ""
	c := rune(l.s[l.pos])

	if unicode.IsLetter(c) {
		s += string(c)
		l.pos++
	}

	for l.pos < len(l.s) {
		c = rune(l.s[l.pos])
		switch {
		case unicode.IsDigit(c) || unicode.IsLetter(c):
			s += string(c)
			l.pos++
		default:
			return s, nil
		}
	}

	return s, nil
}

func (l *lex) getTokenType(s string) int {
	switch s {
	case "var":
		return VAR
	case "def":
		return DEF
//	case "if":
//		return IF
//	case "else":
//		return ELSE
	case "print":
		return TPRINT
	case "return":
		return RETURN
	}
	return ID
}

func (l *lex) Lex(lval *yySymType) int {
	for l.pos < len(l.s) {
		c := rune(l.s[l.pos])

		// >=, <=, &&, ||, !=, ==
		if l.pos < len(l.s)-1 {
			ct := string(c) + string(rune(l.s[l.pos+1]))
			switch ct {
			case ">=":
				l.pos += 2
				return TGEQ
			case "<=":
				l.pos += 2
				return TLEQ
			case "&&":
				l.pos += 2
				return TAND
			case "||":
				l.pos += 2
				return TOR
			case "!=":
				l.pos += 2
				return TNEQ
			case "==":
				l.pos += 2
				return TEQ
			}
		}

		// numbers, ids, tokens
		switch {
		case unicode.IsDigit(c):
			n,e := l.lexNum()
			if e != nil {
				return -1
			}
			lval.e = &IntE{ n }
			return NUMBER
		case unicode.IsLetter(c):
			s,e := l.lexTok()
			if e != nil {
				return -1
			}
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
