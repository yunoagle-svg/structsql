// Table-driven LL(1) predictive parser with an explicit stack: builds the AST, reports syntax errors, recovers on ';' to report several errors.
package main

import "fmt"

/*
KEYWORDS (case-sensitive)
STARTstmt              PROGRAMME
FINALstmt              ;
ENDstmt                FIN
TABLECREATINGstmt      CREATE
TABLEstmt              TABLE
INSERTINGstmt          INSERT
INTOstmt               INTO
VALUESstmt             VALUES
SELECTstmt             SELECT
FROMstmt               FROM
WHEREstmt              WHERE
TYPEINTstmt            INTEGER
TYPEFLOATstmt          REAL
TYPESTRINGstmt         string

PUNCTUATION
OPENEDP                (
CLOSEDP                )
COMMA                  ,
OPENEDB                [
CLOSEDB                ]

OPERATORS
LESS_THAN              <
GREATER_THAN           >
NOT_EQUAL              !=
EQUAL                  =
LESS_THAN_OR_EQUAL     <=
GREATER_THAN_OR_EQUAL  >=
STAR                   *

LITERALS
IDEN                   e.g. Etudiant, moyenne1
NUMBER_INT             e.g. 12, 55, 0
NUMBER_REAL            e.g. 1.5, 14.5, 20.0
I_string               e.g. 'benali'   (whole thing one token, quotes included)

SPECIAL
UNKNOWN_TOKEN          zero value of Kind, never emitted on purpose
EOF                    end of source, appended once, last in slice
LEXICAL_ERROR          unrecognized character, reported but scanning continues

S -> STARTING stmts ENDING

STARTING -> STARTstmt IDEN FINALstmt
ENDING   -> ENDstmt

stmts -> line stmts | EMPTY

line -> CREATING
      | INSERTING
      | QUERIES

--------- create section start -------------------
CREATING -> TABLECREATINGstmt TABLEstmt IDEN OPENEDP CREATINGEXP CLOSEDP FINALstmt

CREATINGEXP -> IDEN TYPE COMMACASE_C

COMMACASE_C -> COMMA IDEN TYPE COMMACASE_C | EMPTY

TYPE -> TYPEINTstmt
      | TYPEFLOATstmt
      | TYPESTRINGstmt OPENEDB NUMBER_INT CLOSEDB

--------- creating section done ------------------

--------- insertion section start ----------------
INSERTING -> INSERTINGstmt INTOstmt IDEN VALUESstmt OPENEDP INSERTIONEXP CLOSEDP FINALstmt

INSERTIONEXP -> LITERALVALUES COMMACASES_I

LITERALVALUES -> NUMBER_INT
               | NUMBER_REAL
               | I_string

COMMACASES_I -> COMMA LITERALVALUES COMMACASES_I | EMPTY

--------- insertion section done -----------------

--------- queries section start -------------------
QUERIES -> SELECTstmt SELECTEXP FROMstmt FROMEXP WHERECASE FINALstmt

SELECTEXP -> STAR
           | IDEN COMMACASE_SE

COMMACASE_SE -> COMMA IDEN COMMACASE_SE | EMPTY

FROMEXP -> IDEN

WHERECASE -> WHEREstmt CONDITON | EMPTY

CONDITON -> CONDITIONEXP OP CONDITIONEXP

CONDITIONEXP -> NUMBER_INT
              | NUMBER_REAL
              | I_string
              | IDEN

OP -> LESS_THAN
    | GREATER_THAN
    | NOT_EQUAL
    | EQUAL
    | LESS_THAN_OR_EQUAL
    | GREATER_THAN_OR_EQUAL

*/
/*
{Literal:PROGRAMME line:0 column:9 TokenType:1}
{Literal:OpTest line:0 column:16 TokenType:27}
{Literal:; line:0 column:16 TokenType:2}
{Literal:SELECT line:1 column:6 TokenType:9}
*/
type Symbol struct {
	ID int
	IsTerminal bool
}
type parser struct {
	tokens []Token
	stack []Symbol
	pos int
	line int
	column int
}


/*
Stack [EOF($), S(axiome)]
chaineOfTokens(from the lexer) [STARTstmt(PROGRAMME){0}, IDEN(Etudiant){1}, FINALstmt(;){2}, ENDstmt(FIN){3}, EOF{4}]
{n} is the position of each token in the chaine

1 -> [EOF, S]     top : S    tok: PROGRAMME  pos : 0

S is non-terminal so we look for table["S"]["STARTstmt"] = [STARTING, stmts, ENDING]
so we pop S
and we push ENDING then stmts then STARTING to the stack
stack becomes [EOF, ENDING, stmts, STARTING] top : STARTING(NT)  tok: PROGRAMME pos : 0


2 -> [EOF, ENDING, stmts, STARTING] top : STARTING(NT)  tok: PROGRAMME pos : 0

STARTING is non-terminal so we look for table["STARTING"]["STARTstmt"] = [STARTstmt, IDEN, FINALstmt]
so we pop STARTING
and push FINALstmt then IDEN then STARTstmt
stack becomes [EOF, ENDING, stmts, FINALstmt, IDEN, STARTstmt] top : STARTstmt tok : PROGRAMME pos : 0


3 -> [EOF, ENDING, stmts, FINALstmt, IDEN, STARTstmt] top : STARTstmt tok : PROGRAMME pos : 0
STARTstmt(PROGRAMME) is terminal so we check if top == tok => STARTstmt(PROGRAMME) == PROGRAMME => true
so we check if top == EOF (stack is empty) if it is we accept
else we pop(top)
and pos++
stack becomes [EOF, ENDING, stmts, FINALstmt, IDEN] top : IDEN tok : IDEN(Etudiant) pos : 1


4 -> [EOF, ENDING, stmts, FINALstmt, IDEN] top : IDEN tok : IDEN(Etudiant) pos : 1
IDEN is terminal so we check if top == tok => IDEN == IDEN(Etudiant) => true
top != EOF so we pop(top)
and pos++
stack becomes [EOF, ENDING, stmts, FINALstmt] top : FINALstmt tok : FINALstmt(;) pos : 2


5 -> [EOF, ENDING, stmts, FINALstmt] top : FINALstmt tok : FINALstmt(;) pos : 2
FINALstmt(;) is terminal so we check if top == tok => FINALstmt(;) == FINALstmt(;) => true
top != EOF so we pop(top)
and pos++
stack becomes [EOF, ENDING, stmts] top : stmts tok : ENDstmt(FIN) pos : 3


6 -> [EOF, ENDING, stmts] top : stmts tok : ENDstmt(FIN) pos : 3
stmts is non-terminal so we look for table["stmts"]["ENDstmt"] = [] (EMPTY)
so we pop stmts
and we push nothing (empty production)
stack becomes [EOF, ENDING] top : ENDING tok : ENDstmt(FIN) pos : 3


7 -> [EOF, ENDING] top : ENDING tok : ENDstmt(FIN) pos : 3
ENDING is non-terminal so we look for table["ENDING"]["ENDstmt"] = [ENDstmt]
so we pop ENDING
and push ENDstmt
stack becomes [EOF, ENDstmt] top : ENDstmt tok : ENDstmt(FIN) pos : 3


8 -> [EOF, ENDstmt] top : ENDstmt tok : ENDstmt(FIN) pos : 3
ENDstmt(FIN) is terminal so we check if top == tok => ENDstmt(FIN) == ENDstmt(FIN) => true
top != EOF so we pop(top)
and pos++
stack becomes [EOF] top : EOF tok : EOF pos : 4


9 -> [EOF] top : EOF tok : EOF pos : 4
EOF is terminal so we check if top == tok => EOF == EOF => true
top == EOF and tok == EOF so we accept 

*/
func (p *parser) Parse() error {
	for len(p.stack) > 0 {
		top := p.stack[len(p.stack)-1]
		tok := p.tokens[p.pos].TokenType

		if isItTerminal(top) {
			if isMatch(Kind(top.ID), tok) {
				// Accept BEFORE popping: top is EOF-terminal and tok is EOF
				if isMatch(Kind(top.ID), EOF) && tok == EOF {
					return nil
				}
				p.pop()
				p.pos++
			} else {
				return fmt.Errorf("expected %v, got %v at line %d",
					Kind(top.ID).String(), tok.String(), p.tokens[p.pos].line)
			}
		} else {
			rhs, err := lookFor(tok, Nonterminal(top.ID))
			if err != nil {
				return err
			}
			p.pop()
			p.push(rhs)
		}
	}
	return fmt.Errorf("parse ended without accepting")
}

//--------------- HELPER FUNCTIONS


func createParser(tokens []Token) *parser {
	return &parser{
		tokens: tokens,
		pos: 0,
		stack: []Symbol{T(EOF), NT(S)},
	}
}

func NT(n Nonterminal) Symbol {
	return Symbol{ID: int(n), IsTerminal: false}
}

func T(k Kind) Symbol {
	 return Symbol{ID: int(k), IsTerminal: true}
}

func lookFor(tr Kind, nt Nonterminal) ([]Symbol, error) {
    if production, exist := table[nt][tr]; exist {
        return production, nil
    }
    return nil, fmt.Errorf("no production found for nonterminal %v and token %v", nt.String(), tr.String())
}

func (p *parser) push(s []Symbol) {
    for i := len(s) - 1; i >= 0; i-- {
        p.stack = append(p.stack, s[i])
    }
}
func (p *parser) pop() {
	p.stack = p.stack[:len(p.stack)-1]
}
func isItTerminal(s Symbol) bool {
	return s.IsTerminal
}
func isMatch(i Kind, ii Kind) bool {
	return i == ii
}

func main() {
	source := `PROGRAMME x;`

	l := createLexer(source)
	t, err0 := l.splitIntoTokens()
	if err0 != nil {
		fmt.Println("lexer error:", err0)
		return
	}

	p := createParser(t)
	err := p.Parse()
	if err != nil {
		fmt.Println("parser error:", err)
		return
	}

	fmt.Println("woah — accepted")
}

