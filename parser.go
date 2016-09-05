//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:3
//line parser.y:8
type yySymType struct {
	yys int
	str string
	s   S
	e   E
	fd  *FDefS
	sl  *SL
	idl []string
	el  []E
}

const NUMBER = 57346
const ID = 57347
const RETURN = 57348
const DEF = 57349
const VAR = 57350
const TPRINT = 57351
const UMINUS = 57352
const TLEQ = 57353
const TGEQ = 57354
const TEQ = 57355
const TNEQ = 57356
const TAND = 57357
const TOR = 57358

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
	"'='",
	"';'",
	"'('",
	"')'",
	"'{'",
	"'}'",
	"','",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:135

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 36
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 190

var yyAct = [...]int{

	67, 68, 11, 34, 35, 36, 37, 38, 40, 44,
	39, 43, 41, 42, 45, 46, 20, 12, 70, 49,
	25, 30, 71, 10, 7, 32, 45, 46, 47, 48,
	42, 45, 46, 51, 31, 23, 53, 54, 55, 56,
	57, 58, 59, 60, 61, 62, 63, 64, 65, 34,
	35, 36, 37, 38, 40, 44, 39, 43, 41, 42,
	45, 46, 4, 22, 9, 66, 41, 42, 45, 46,
	14, 6, 72, 34, 35, 36, 37, 38, 40, 44,
	39, 43, 41, 42, 45, 46, 21, 69, 34, 35,
	36, 37, 38, 40, 44, 39, 43, 41, 42, 45,
	46, 1, 52, 34, 35, 36, 37, 38, 40, 44,
	39, 43, 41, 42, 45, 46, 3, 50, 34, 35,
	36, 37, 38, 40, 44, 39, 43, 41, 42, 45,
	46, 15, 33, 34, 35, 36, 37, 38, 40, 44,
	39, 43, 41, 42, 45, 46, 36, 37, 38, 40,
	44, 39, 43, 41, 42, 45, 46, 40, 44, 39,
	43, 41, 42, 45, 46, 29, 28, 29, 28, 8,
	2, 0, 0, 27, 5, 27, 17, 18, 0, 16,
	19, 13, 0, 0, 0, 0, 24, 26, 0, 26,
}
var yyPact = [...]int{

	55, -1000, -1000, 55, 66, -1000, -2, 59, -4, -28,
	-11, 59, 171, -1000, -13, 171, 58, 11, 161, 163,
	-1000, -1000, 10, 163, -1000, 107, 163, 163, -7, -1000,
	92, 163, 77, -1000, 163, 163, 163, 163, 163, 163,
	163, 163, 163, 163, 163, 163, 163, 38, 122, 163,
	-1000, 62, -1000, 133, 133, 141, 141, 141, 46, 46,
	9, 4, 46, 46, -1000, -1000, -1000, -9, -8, -1000,
	-1000, 163, -1000,
}
var yyPgo = [...]int{

	0, 1, 131, 116, 101, 170, 70, 169, 0,
}
var yyR1 = [...]int{

	0, 4, 5, 5, 6, 6, 2, 2, 2, 2,
	2, 3, 7, 7, 7, 8, 8, 8, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1,
}
var yyR2 = [...]int{

	0, 1, 0, 2, 0, 2, 5, 4, 2, 3,
	3, 8, 0, 1, 3, 0, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 1, 4, 1,
}
var yyChk = [...]int{

	-1000, -4, -5, -3, 7, -5, 5, 26, -7, 5,
	27, 30, 28, -7, -6, -2, 8, 5, 6, 9,
	29, -6, 5, 24, 25, -1, 26, 12, 5, 4,
	-1, 24, -1, 25, 11, 12, 13, 14, 15, 18,
	16, 20, 21, 19, 17, 22, 23, -1, -1, 26,
	25, -1, 25, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, 27, -8, -1, 25,
	27, 30, -8,
}
var yyDef = [...]int{

	2, -2, 1, 2, 0, 3, 0, 12, 0, 13,
	0, 12, 4, 14, 0, 4, 0, 0, 0, 0,
	11, 5, 0, 0, 8, 0, 0, 0, 33, 35,
	0, 0, 0, 9, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 32, 15,
	10, 0, 7, 18, 19, 20, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 30, 31, 0, 16, 6,
	34, 15, 17,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 15, 3, 3,
	26, 27, 13, 11, 30, 12, 3, 14, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 25,
	16, 24, 18, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 28, 3, 29,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 17,
	19, 20, 21, 22, 23,
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
		//line parser.y:43
		{
			yylex.(*lex).result = yyVAL.sl
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:48
		{
			yyVAL.sl = &SL{}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:50
		{
			yyVAL.sl = &SL{ss: append([]S{yyDollar[1].fd}, yyDollar[2].sl.ss...)}
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:55
		{
			yyVAL.sl = &SL{}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:57
		{
			yyVAL.sl = &SL{ss: append([]S{yyDollar[1].s}, yyDollar[2].sl.ss...)}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:62
		{
			yyVAL.s = &VDefS{name: yyDollar[2].str, rhs: yyDollar[4].e}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:64
		{
			yyVAL.s = &AssignS{name: yyDollar[1].str, rhs: yyDollar[3].e}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:66
		{
			yyVAL.s = &RetS{}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:68
		{
			yyVAL.s = &RetS{rhs: yyDollar[2].e}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:70
		{
			yyVAL.s = &PrintS{e: yyDollar[2].e}
		}
	case 11:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:75
		{
			yyVAL.fd = &FDefS{name: yyDollar[2].str, args: yyDollar[4].idl, body: yyDollar[7].sl}
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:80
		{
			yyVAL.idl = []string{}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:82
		{
			yyVAL.idl = []string{yyDollar[1].str}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:84
		{
			yyVAL.idl = append([]string{yyDollar[1].str}, yyDollar[3].idl...)
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:89
		{
			yyVAL.el = []E{}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:91
		{
			yyVAL.el = []E{yyDollar[1].e}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:93
		{
			yyVAL.el = append([]E{yyDollar[1].e}, yyDollar[3].el...)
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:98
		{
			yyVAL.e = &BinOpE{ADD, yyDollar[1].e, yyDollar[3].e}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:100
		{
			yyVAL.e = &BinOpE{SUB, yyDollar[1].e, yyDollar[3].e}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:102
		{
			yyVAL.e = &BinOpE{MUL, yyDollar[1].e, yyDollar[3].e}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:104
		{
			yyVAL.e = &BinOpE{DIV, yyDollar[1].e, yyDollar[3].e}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:106
		{
			yyVAL.e = &BinOpE{MOD, yyDollar[1].e, yyDollar[3].e}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:108
		{
			yyVAL.e = &BinOpE{GT, yyDollar[1].e, yyDollar[3].e}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:110
		{
			yyVAL.e = &BinOpE{LT, yyDollar[1].e, yyDollar[3].e}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:112
		{
			yyVAL.e = &BinOpE{EQ, yyDollar[1].e, yyDollar[3].e}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:114
		{
			yyVAL.e = &BinOpE{NEQ, yyDollar[1].e, yyDollar[3].e}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:116
		{
			yyVAL.e = &BinOpE{GEQ, yyDollar[1].e, yyDollar[3].e}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:118
		{
			yyVAL.e = &BinOpE{LEQ, yyDollar[1].e, yyDollar[3].e}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:120
		{
			yyVAL.e = &BinOpE{AND, yyDollar[1].e, yyDollar[3].e}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:122
		{
			yyVAL.e = &BinOpE{OR, yyDollar[1].e, yyDollar[3].e}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:124
		{
			yyVAL.e = yyDollar[2].e
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:126
		{
			yyVAL.e = &BinOpE{MUL, &IntE{-1}, yyDollar[2].e}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:128
		{
			yyVAL.e = &IdE{yyDollar[1].str}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:130
		{
			yyVAL.e = &CallE{name: yyDollar[1].str, args: yyDollar[3].el}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:132
		{
			yyVAL.e = yyDollar[1].e
		}
	}
	goto yystack /* stack new state and value */
}
