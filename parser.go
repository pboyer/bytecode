//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:3
import "fmt"

//line parser.y:10
type yySymType struct {
	yys int
	str string
	s   S
	e   E
	fd  *FDefS
	sl  *BlockS
	idl []string
	el  []E
}

const NUMBER = 57346
const ID = 57347
const RETURN = 57348
const DEF = 57349
const VAR = 57350
const TPRINT = 57351
const IF = 57352
const ELSE = 57353
const UMINUS = 57354
const TLEQ = 57355
const TGEQ = 57356
const TEQ = 57357
const TNEQ = 57358
const TAND = 57359
const TOR = 57360
const THEN = 57361

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"ID",
	"RETURN",
	"DEF",
	"VAR",
	"TPRINT",
	"IF",
	"ELSE",
	"UMINUS",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'<'",
	"TLEQ",
	"'>'",
	"TGEQ",
	"TEQ",
	"TNEQ",
	"TAND",
	"TOR",
	"THEN",
	"'='",
	"';'",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"','",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:146

func parse(prog string) (*BlockS, error) {
	l := &lex{s: prog}
	r := yyParse(l)
	if r != 0 {
		return nil, fmt.Errorf("Unknown parser error encountered")
	}

	return l.result, nil
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 39
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 219

var yyAct = [...]int{

	15, 73, 74, 11, 39, 40, 41, 42, 43, 45,
	49, 44, 48, 46, 47, 50, 51, 77, 10, 54,
	34, 7, 28, 33, 78, 37, 22, 12, 36, 47,
	50, 51, 52, 53, 35, 25, 81, 56, 57, 50,
	51, 4, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 39, 40, 41, 42, 43,
	45, 49, 44, 48, 46, 47, 50, 51, 24, 32,
	31, 9, 6, 2, 75, 1, 79, 5, 3, 30,
	80, 0, 82, 39, 40, 41, 42, 43, 45, 49,
	44, 48, 46, 47, 50, 51, 29, 46, 47, 50,
	51, 0, 72, 39, 40, 41, 42, 43, 45, 49,
	44, 48, 46, 47, 50, 51, 0, 0, 76, 39,
	40, 41, 42, 43, 45, 49, 44, 48, 46, 47,
	50, 51, 0, 0, 58, 39, 40, 41, 42, 43,
	45, 49, 44, 48, 46, 47, 50, 51, 0, 0,
	55, 39, 40, 41, 42, 43, 45, 49, 44, 48,
	46, 47, 50, 51, 0, 0, 38, 39, 40, 41,
	42, 43, 45, 49, 44, 48, 46, 47, 50, 51,
	41, 42, 43, 45, 49, 44, 48, 46, 47, 50,
	51, 32, 31, 17, 19, 14, 16, 20, 21, 0,
	8, 30, 45, 49, 44, 48, 46, 47, 50, 51,
	0, 23, 13, 0, 26, 27, 0, 18, 29,
}
var yyPact = [...]int{

	34, -1000, -1000, 34, 67, -1000, -10, 66, -14, -30,
	-2, 66, 188, -1000, -4, 188, 63, 8, 188, 187,
	65, -11, -1000, -1000, 7, 65, -5, -1000, 138, 65,
	65, -12, -1000, 122, 65, 65, 106, -1000, -1000, 65,
	65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
	65, 65, 70, 154, 65, -1000, 42, 90, -1000, 165,
	165, 184, 184, 184, 75, 75, 6, 15, 75, 75,
	-1000, -1000, -1000, -15, -9, 188, -1000, -1000, 65, 25,
	-1000, 188, -1000,
}
var yyPgo = [...]int{

	0, 2, 0, 78, 75, 73, 195, 200, 1,
}
var yyR1 = [...]int{

	0, 4, 5, 5, 6, 6, 2, 2, 2, 2,
	2, 2, 2, 2, 3, 7, 7, 7, 8, 8,
	8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1,
}
var yyR2 = [...]int{

	0, 1, 0, 2, 0, 2, 5, 4, 3, 2,
	3, 3, 5, 7, 8, 0, 1, 3, 0, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 1, 4, 1,
}
var yyChk = [...]int{

	-1000, -4, -5, -3, 7, -5, 5, 31, -7, 5,
	32, 33, 29, -7, -6, -2, 8, 5, 29, 6,
	9, 10, 30, -6, 5, 27, -6, 28, -1, 31,
	14, 5, 4, -1, 31, 27, -1, 30, 28, 13,
	14, 15, 16, 17, 20, 18, 22, 23, 21, 19,
	24, 25, -1, -1, 31, 28, -1, -1, 28, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, 32, -8, -1, 32, 28, 32, 33, -2,
	-8, 11, -2,
}
var yyDef = [...]int{

	2, -2, 1, 2, 0, 3, 0, 15, 0, 16,
	0, 15, 4, 17, 0, 4, 0, 0, 4, 0,
	0, 0, 14, 5, 0, 0, 0, 9, 0, 0,
	0, 36, 38, 0, 0, 0, 0, 8, 10, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 35, 18, 11, 0, 0, 7, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 0, 19, 0, 6, 37, 18, 12,
	20, 0, 13,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 17, 3, 3,
	31, 32, 15, 13, 33, 14, 3, 16, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 28,
	18, 27, 20, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 29, 3, 30,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 19, 21, 22, 23, 24, 25, 26,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:48
		{
			yylex.(*lex).result = yyVAL.sl
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:53
		{
			yyVAL.sl = &BlockS{}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:55
		{
			yyVAL.sl = &BlockS{list: append([]S{yyDollar[1].fd}, yyDollar[2].sl.list...)}
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:60
		{
			yyVAL.sl = &BlockS{}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:62
		{
			yyVAL.sl = &BlockS{list: append([]S{yyDollar[1].s}, yyDollar[2].sl.list...)}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:67
		{
			yyVAL.s = &VDefS{name: yyDollar[2].str, rhs: yyDollar[4].e}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:69
		{
			yyVAL.s = &AssignS{name: yyDollar[1].str, rhs: yyDollar[3].e}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:71
		{
			yyVAL.s = yyDollar[2].sl
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:73
		{
			yyVAL.s = &RetS{}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:75
		{
			yyVAL.s = &RetS{rhs: yyDollar[2].e}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:77
		{
			yyVAL.s = &PrintS{e: yyDollar[2].e}
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:79
		{
			yyVAL.s = &IfS{test: yyDollar[3].e, tb: yyDollar[5].s}
		}
	case 13:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:81
		{
			yyVAL.s = &IfS{test: yyDollar[3].e, tb: yyDollar[5].s, fb: yyDollar[7].s}
		}
	case 14:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:86
		{
			yyVAL.fd = &FDefS{name: yyDollar[2].str, args: yyDollar[4].idl, body: yyDollar[7].sl}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:91
		{
			yyVAL.idl = []string{}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:93
		{
			yyVAL.idl = []string{yyDollar[1].str}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:95
		{
			yyVAL.idl = append([]string{yyDollar[1].str}, yyDollar[3].idl...)
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:100
		{
			yyVAL.el = []E{}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:102
		{
			yyVAL.el = []E{yyDollar[1].e}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:104
		{
			yyVAL.el = append([]E{yyDollar[1].e}, yyDollar[3].el...)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:109
		{
			yyVAL.e = &BinOpE{ADD, yyDollar[1].e, yyDollar[3].e}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:111
		{
			yyVAL.e = &BinOpE{SUB, yyDollar[1].e, yyDollar[3].e}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:113
		{
			yyVAL.e = &BinOpE{MUL, yyDollar[1].e, yyDollar[3].e}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:115
		{
			yyVAL.e = &BinOpE{DIV, yyDollar[1].e, yyDollar[3].e}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:117
		{
			yyVAL.e = &BinOpE{MOD, yyDollar[1].e, yyDollar[3].e}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:119
		{
			yyVAL.e = &BinOpE{GT, yyDollar[1].e, yyDollar[3].e}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:121
		{
			yyVAL.e = &BinOpE{LT, yyDollar[1].e, yyDollar[3].e}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:123
		{
			yyVAL.e = &BinOpE{EQ, yyDollar[1].e, yyDollar[3].e}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:125
		{
			yyVAL.e = &BinOpE{NEQ, yyDollar[1].e, yyDollar[3].e}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:127
		{
			yyVAL.e = &BinOpE{GEQ, yyDollar[1].e, yyDollar[3].e}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:129
		{
			yyVAL.e = &BinOpE{LEQ, yyDollar[1].e, yyDollar[3].e}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:131
		{
			yyVAL.e = &BinOpE{AND, yyDollar[1].e, yyDollar[3].e}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:133
		{
			yyVAL.e = &BinOpE{OR, yyDollar[1].e, yyDollar[3].e}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:135
		{
			yyVAL.e = yyDollar[2].e
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:137
		{
			yyVAL.e = &BinOpE{MUL, &IntE{-1}, yyDollar[2].e}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:139
		{
			yyVAL.e = &IdE{yyDollar[1].str}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:141
		{
			yyVAL.e = &CallE{name: yyDollar[1].str, args: yyDollar[3].el}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:143
		{
			yyVAL.e = yyDollar[1].e
		}
	}
	goto yystack /* stack new state and value */
}
