package main

// Nonterminal identifies a left-hand-side grammar symbol.
type Nonterminal int

const (
	S Nonterminal = iota
	STARTING
	ENDING
	stmts
	line
	CREATING
	CREATINGEXP
	COMMACASE_C
	TYPE
	INSERTING
	INSERTIONEXP
	LITERALVALUES
	COMMACASES_I
	QUERIES
	SELECTEXP
	COMMACASE_SE
	FROMEXP
	WHERECASE
	CONDITON
	CONDITIONEXP
	OP
	DOLAR
)

