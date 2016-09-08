%{

package main

import "fmt"

%}

// All the different syntax types
%union{
	str string
	s S
	e E
	fd *FDefS
	sl *BlockS
	idl []string
	el []E
}

// Assignment of types to non-terminals
%type <e> e
%type <s> s
%type <fd> fd
%type <sl> prog fdl sl
%type <idl> idl
%type <el> el

// Assignment of types to terminals
%token <e> NUMBER
%token <str> ID RETURN DEF VAR TPRINT IF ELSE

// Precedence
%left UMINUS
%left '+'  '-'
%left '*'  '/'  '%'
%left '<'  TLEQ '>' TGEQ
%left TEQ
%left TNEQ
%left TAND TOR

%nonassoc THEN
%nonassoc ELSE

%%

prog
	: fdl
	{ yylex.(*lex).result = $$ }
	;

fdl
	: /* empty */
	{ $$ = &BlockS{} }
	| fd fdl
	{ $$ = &BlockS{ list : append([]S{ $1 }, $2.list...) } }
	;

sl
	: /* empty */
	{ $$ = &BlockS{} }
	| s sl
	{ $$ = &BlockS{ list : append([]S{ $1 }, $2.list...) } }
	;

s
	: VAR ID '=' e ';'
	{ $$ = &VDefS{ name : $2, rhs : $4 } }
	| ID '=' e ';'
	{ $$ = &AssignS{ name : $1, rhs : $3 } }
	| '{' sl '}'
	{ $$ = $2 }
	| RETURN ';'
	{ $$ = &RetS{} }
	| RETURN e ';'
	{ $$ = &RetS{ rhs : $2 } }
	| TPRINT e ';'
	{ $$ = &PrintS{ e : $2 } }
	| IF '(' e ')' s     %prec THEN
	{ $$ = &IfS{ test : $3, tb : $5 } }
	| IF '(' e ')' s ELSE s
	{ $$ = &IfS{ test : $3, tb : $5, fb : $7 } }
	;

fd
	: DEF ID '(' idl ')' '{' sl '}'
	{ $$ = &FDefS{ name : $2, args : $4, body : $7} }
	;

idl
	: /* empty */
	{ $$ = []string{} }
	| ID
	{ $$ = []string{ $1 } }
	| ID ',' idl
	{ $$ = append([]string{ $1 }, $3...) }
	;

el
	: /* empty */
	{ $$ = []E{} }
	| e
	{ $$ = []E{ $1 } }
	| e ',' el
	{ $$ = append([]E{ $1 }, $3...) }
	;

e
	: e '+' e
	{ $$  =  &BinOpE{ ADD, $1, $3 } }
	| e '-' e
	{ $$  =  &BinOpE{ SUB, $1, $3 } }
	| e '*' e
	{ $$  =  &BinOpE{ MUL, $1, $3 } }
	| e '/' e
	{ $$  =  &BinOpE{ DIV, $1, $3 } }
	| e '%' e
	{ $$  =  &BinOpE{ MOD, $1, $3 } }
	| e '>' e
	{ $$  =  &BinOpE{ GT, $1, $3 } }
	| e '<' e
	{ $$  =  &BinOpE{ LT, $1, $3 } }
	| e TEQ e
	{ $$  =  &BinOpE{ EQ, $1, $3 } }
	| e TNEQ e
	{ $$  =  &BinOpE{ NEQ, $1, $3 } }
	| e TGEQ e
	{ $$  =  &BinOpE{ GEQ, $1, $3 } }
	| e TLEQ e
	{ $$  =  &BinOpE{ LEQ, $1, $3 } }
	| e TAND e
	{ $$  =  &BinOpE{ AND, $1, $3 } }
	| e TOR e
	{ $$  =  &BinOpE{ OR, $1, $3 } }
	| '(' e ')'
	{ $$  =  $2 }
	| '-'  e     %prec  UMINUS
	{ $$  = &BinOpE{ MUL, &IntE{ -1 }, $2 } }
	| ID
	{ $$  =  &IdE{ $1 } }
	| ID '(' el ')'
	{ $$  =  &CallE{ name : $1, args : $3 } }
	| NUMBER
	{ $$ = $1 }
	;

%%

func parse(prog string) (*BlockS, error) {
	l := &lex{s: prog}
	r := yyParse(l)
	if r != 0 {
		return nil, fmt.Errorf("Unknown parser error encountered")
	}

	return l.result, nil
}