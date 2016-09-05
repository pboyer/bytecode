%{

package main

%}

// All the different syntax types
%union{
	str string
	s S
	e E
	fd *FDefS
	sl *SL
	idl []string
}

// Assignment of types to non-terminals
%type <e> e
%type <s> s
%type <fd> fd
%type <sl> prog fdl sl
%type <idl> idl

// Assignment of types to terminals
%token <e> NUMBER
%token <str> ID RETURN DEF VAR PRINT

// Arithmetic precedence
%left '+'  '-'
%left '*'  '/'  '%'

%%

prog
	: fdl
	{ parserlex.(*lex).result = $$ }
	;

fdl
	: /* empty */
	{ $$ = &SL{} }
	| fd fdl
	{ $$ = &SL{ ss : append([]S{ $1 }, $2.ss...) } }
	;

sl
	: /* empty */
	{ $$ = &SL{} }
	| s sl
	{ $$ = &SL{ ss : append([]S{ $1 }, $2.ss...) } }
	;

s
	: VAR ID '=' e ';'
	{ $$ = &VDefS{ name : $2, rhs : $4 } }
	| ID '=' e ';'
	{ $$ = &AssignS{ name : $1, rhs : $3 } }
	| RETURN e ';'
	{ $$ = &RetS{ rhs : $2 } }
	| PRINT e ';'
	{ $$ = &PrintS{ e : $2 } }
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

e
	: e '+' e
	{ $$  =  &BinOpE{ ADD, $1, $3 } }
	| e '-' e
	{ $$  =  &BinOpE{ SUB, $1, $3 } }
	| NUMBER
	{ $$ = $1 }
	;

%%
