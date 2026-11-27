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
