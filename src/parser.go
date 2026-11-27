// Table-driven LL(1) predictive parser with an explicit stack: builds the AST, reports syntax errors, recovers on ';' to report several errors.
package main

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

S -> STARTING stmts ENDING | empty ;

STARTING -> STARTstmt IDEN FINALstmt;
ENDING   -> ENDstmt EOF

stmts ->
      line
	  line stmts
	  ;
line ->
      CREATING
	  INSERTING
	  QUERIES
	  EMPTY
	;
--------- create section start  -------------------
CREATING -> CREATEstmt TABLEstmt IDEN OPENEDP CREATINGEXP CLOSEDP FINALstmt
CREATINGEXP ->
      IDEN TYPE COMMACASE
	;
COMMACASE ->
      COMMA IDEN TYPE
	  EMPTY
	;
TYPE ->
      TYPEINTstmt
	  TYPEFLOATstmt
	  TYPESTRINGstmt
	;
--------- creating section DONE ------------------

*/

type parser struct {
	tokens []Token
	pos int
}
/*type grammar_elements struct {
	terminals []Token
	non_terminals []
}*/

func (p parser) nextToken() Kind {
	return p.tokens[p.pos + 1].TokenType
}

func (p parser) currentToken() Kind {
	return p.tokens[p.pos].TokenType
}

func createParser(tokens []Token) *parser {
	return &parser{
		tokens: tokens,
		pos: 0,
	}
}
