//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:3
//line parser.y:8
type parserSymType struct {
	yys int
	str string
	s   S
	e   E
	fd  *FDefS
	sl  *SL
	idl []string
}

const NUMBER = 57346
const ID = 57347
const RETURN = 57348
const DEF = 57349
const VAR = 57350

var parserToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"ID",
	"RETURN",
	"DEF",
	"VAR",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'='",
	"';'",
	"'('",
	"')'",
	"'{'",
	"'}'",
	"','",
}
var parserStatenames = [...]string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line parser.y:78

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const parserNprod = 14
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 37

var parserAct = [...]int{

	22, 10, 18, 11, 27, 28, 27, 28, 27, 28,
	33, 9, 30, 7, 26, 13, 6, 24, 21, 3,
	16, 17, 25, 15, 12, 29, 20, 8, 31, 32,
	19, 5, 23, 1, 2, 14, 4,
}
var parserPact = [...]int{

	12, -1000, 12, 26, -1000, 0, 22, -6, -19, -15,
	22, 15, -1000, -17, 15, 21, 4, 28, -1000, -1000,
	3, 28, -1, -1000, 28, -3, -1000, 28, 28, -5,
	-1000, -1000, -1000, -1000,
}
var parserPgo = [...]int{

	0, 0, 35, 34, 33, 15, 13,
}
var parserR1 = [...]int{

	0, 4, 4, 5, 5, 2, 2, 2, 3, 6,
	6, 1, 1, 1,
}
var parserR2 = [...]int{

	0, 0, 2, 0, 2, 5, 4, 3, 8, 1,
	3, 3, 3, 1,
}
var parserChk = [...]int{

	-1000, -4, -3, 7, -4, 5, 16, -6, 5, 17,
	20, 18, -6, -5, -2, 8, 5, 6, 19, -5,
	5, 14, -1, 4, 14, -1, 15, 9, 10, -1,
	15, -1, -1, 15,
}
var parserDef = [...]int{

	1, -2, 1, 0, 2, 0, 0, 0, 9, 0,
	0, 3, 10, 0, 3, 0, 0, 0, 8, 4,
	0, 0, 0, 13, 0, 0, 7, 0, 0, 0,
	6, 11, 12, 5,
}
var parserTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 13, 3, 3,
	16, 17, 11, 9, 20, 10, 3, 12, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 15,
	3, 14, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 18, 3, 19,
}
var parserTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8,
}
var parserTok3 = [...]int{
	0,
}

var parserErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	parserDebug        = 0
	parserErrorVerbose = false
)

type parserLexer interface {
	Lex(lval *parserSymType) int
	Error(s string)
}

type parserParser interface {
	Parse(parserLexer) int
	Lookahead() int
}

type parserParserImpl struct {
	lookahead func() int
}

func (p *parserParserImpl) Lookahead() int {
	return p.lookahead()
}

func parserNewParser() parserParser {
	p := &parserParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const parserFlag = -1000

func parserTokname(c int) string {
	if c >= 1 && c-1 < len(parserToknames) {
		if parserToknames[c-1] != "" {
			return parserToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func parserStatname(s int) string {
	if s >= 0 && s < len(parserStatenames) {
		if parserStatenames[s] != "" {
			return parserStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func parserErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !parserErrorVerbose {
		return "syntax error"
	}

	for _, e := range parserErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + parserTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := parserPact[state]
	for tok := TOKSTART; tok-1 < len(parserToknames); tok++ {
		if n := base + tok; n >= 0 && n < parserLast && parserChk[parserAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if parserDef[state] == -2 {
		i := 0
		for parserExca[i] != -1 || parserExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; parserExca[i] >= 0; i += 2 {
			tok := parserExca[i]
			if tok < TOKSTART || parserExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if parserExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += parserTokname(tok)
	}
	return res
}

func parserlex1(lex parserLexer, lval *parserSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = parserTok1[0]
		goto out
	}
	if char < len(parserTok1) {
		token = parserTok1[char]
		goto out
	}
	if char >= parserPrivate {
		if char < parserPrivate+len(parserTok2) {
			token = parserTok2[char-parserPrivate]
			goto out
		}
	}
	for i := 0; i < len(parserTok3); i += 2 {
		token = parserTok3[i+0]
		if token == char {
			token = parserTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = parserTok2[1] /* unknown char */
	}
	if parserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", parserTokname(token), uint(char))
	}
	return char, token
}

func parserParse(parserlex parserLexer) int {
	return parserNewParser().Parse(parserlex)
}

func (parserrcvr *parserParserImpl) Parse(parserlex parserLexer) int {
	var parsern int
	var parserlval parserSymType
	var parserVAL parserSymType
	var parserDollar []parserSymType
	_ = parserDollar // silence set and not used
	parserS := make([]parserSymType, parserMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	parserstate := 0
	parserchar := -1
	parsertoken := -1 // parserchar translated into internal numbering
	parserrcvr.lookahead = func() int { return parserchar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		parserstate = -1
		parserchar = -1
		parsertoken = -1
	}()
	parserp := -1
	goto parserstack

ret0:
	return 0

ret1:
	return 1

parserstack:
	/* put a state and value onto the stack */
	if parserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", parserTokname(parsertoken), parserStatname(parserstate))
	}

	parserp++
	if parserp >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserS[parserp] = parserVAL
	parserS[parserp].yys = parserstate

parsernewstate:
	parsern = parserPact[parserstate]
	if parsern <= parserFlag {
		goto parserdefault /* simple state */
	}
	if parserchar < 0 {
		parserchar, parsertoken = parserlex1(parserlex, &parserlval)
	}
	parsern += parsertoken
	if parsern < 0 || parsern >= parserLast {
		goto parserdefault
	}
	parsern = parserAct[parsern]
	if parserChk[parsern] == parsertoken { /* valid shift */
		parserchar = -1
		parsertoken = -1
		parserVAL = parserlval
		parserstate = parsern
		if Errflag > 0 {
			Errflag--
		}
		goto parserstack
	}

parserdefault:
	/* default state action */
	parsern = parserDef[parserstate]
	if parsern == -2 {
		if parserchar < 0 {
			parserchar, parsertoken = parserlex1(parserlex, &parserlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if parserExca[xi+0] == -1 && parserExca[xi+1] == parserstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			parsern = parserExca[xi+0]
			if parsern < 0 || parsern == parsertoken {
				break
			}
		}
		parsern = parserExca[xi+1]
		if parsern < 0 {
			goto ret0
		}
	}
	if parsern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			parserlex.Error(parserErrorMessage(parserstate, parsertoken))
			Nerrs++
			if parserDebug >= 1 {
				__yyfmt__.Printf("%s", parserStatname(parserstate))
				__yyfmt__.Printf(" saw %s\n", parserTokname(parsertoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for parserp >= 0 {
				parsern = parserPact[parserS[parserp].yys] + parserErrCode
				if parsern >= 0 && parsern < parserLast {
					parserstate = parserAct[parsern] /* simulate a shift of "error" */
					if parserChk[parserstate] == parserErrCode {
						goto parserstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if parserDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", parserS[parserp].yys)
				}
				parserp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if parserDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", parserTokname(parsertoken))
			}
			if parsertoken == parserEofCode {
				goto ret1
			}
			parserchar = -1
			parsertoken = -1
			goto parsernewstate /* try again in the same state */
		}
	}

	/* reduction by production parsern */
	if parserDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", parsern, parserStatname(parserstate))
	}

	parsernt := parsern
	parserpt := parserp
	_ = parserpt // guard against "declared and not used"

	parserp -= parserR2[parsern]
	// parserp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if parserp+1 >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserVAL = parserS[parserp+1]

	/* consult goto table to find next state */
	parsern = parserR1[parsern]
	parserg := parserPgo[parsern]
	parserj := parserg + parserS[parserp].yys + 1

	if parserj >= parserLast {
		parserstate = parserAct[parserg]
	} else {
		parserstate = parserAct[parserj]
		if parserChk[parserstate] != -parsern {
			parserstate = parserAct[parserg]
		}
	}
	// dummy call; replaced with literal code
	switch parsernt {

	case 1:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line parser.y:36
		{
			parserVAL.sl = &SL{}
		}
	case 2:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:38
		{
			parserVAL.sl = &SL{ss: append([]S{parserDollar[1].fd}, parserDollar[2].sl.ss...)}
		}
	case 3:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line parser.y:43
		{
			parserVAL.sl = &SL{}
		}
	case 4:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line parser.y:45
		{
			parserVAL.sl = &SL{ss: append([]S{parserDollar[1].s}, parserDollar[2].sl.ss...)}
		}
	case 5:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line parser.y:50
		{
			parserVAL.s = &VDefS{name: parserDollar[2].str, rhs: parserDollar[4].e}
		}
	case 6:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line parser.y:52
		{
			parserVAL.s = &AssignS{name: parserDollar[1].str, rhs: parserDollar[3].e}
		}
	case 7:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:54
		{
			parserVAL.s = &RetS{rhs: parserDollar[2].e}
		}
	case 8:
		parserDollar = parserS[parserpt-8 : parserpt+1]
		//line parser.y:59
		{
			parserVAL.fd = &FDefS{name: parserDollar[2].str, args: parserDollar[4].idl, body: parserDollar[7].sl}
		}
	case 9:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:64
		{
			parserVAL.idl = []string{parserDollar[1].str}
		}
	case 10:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:66
		{
			parserVAL.idl = append([]string{parserDollar[1].str}, parserDollar[3].idl...)
		}
	case 11:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:71
		{
			parserVAL.e = &BinOpE{ADD, parserDollar[1].e, parserDollar[3].e}
		}
	case 12:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line parser.y:73
		{
			parserVAL.e = &BinOpE{SUB, parserDollar[1].e, parserDollar[3].e}
		}
	case 13:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line parser.y:75
		{
			parserVAL.e = parserDollar[1].e
		}
	}
	goto parserstack /* stack new state and value */
}
