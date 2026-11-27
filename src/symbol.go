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
func (n Nonterminal) String() string {
	switch n {
	case S:
		return "S"
	case STARTING:
		return "STARTING"
	case ENDING:
		return "ENDING"
	case stmts:
		return "stmts"
	case line:
		return "line"
	case CREATING:
		return "CREATING"
	case CREATINGEXP:
		return "CREATINGEXP"
	case COMMACASE_C:
		return "COMMACASE_C"
	case TYPE:
		return "TYPE"
	case INSERTING:
		return "INSERTING"
	case INSERTIONEXP:
		return "INSERTIONEXP"
	case LITERALVALUES:
		return "LITERALVALUES"
	case COMMACASES_I:
		return "COMMACASES_I"
	case QUERIES:
		return "QUERIES"
	case SELECTEXP:
		return "SELECTEXP"
	case COMMACASE_SE:
		return "COMMACASE_SE"
	case FROMEXP:
		return "FROMEXP"
	case WHERECASE:
		return "WHERECASE"
	case CONDITON:
		return "CONDITON"
	case CONDITIONEXP:
		return "CONDITIONEXP"
	case OP:
		return "OP"
	case DOLAR:
		return "DOLAR"
	default:
		return "UNKNOWN_NONTERMINAL"
	}
}

