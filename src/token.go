package main

import (
	"fmt"
)

/*

TOKENS

KEYWORDS  :  (
                    PROGRAMME (STARTstmt)          HOW THE PROGRAM SHOULD START WITH
					; (FINALstmt)                  FINAL
					FIN (ENDstmt)                  HOW THE PROGRAM SHOULD END WITH
					INTEGER (TYPEINTstmt)          a whole number
					REAL  (TYPEFLOATstmt)          a floating-point number
					string (TYPESTRINGstmt)        a character string
					CREATE (TABLECREATINGstmt)     for creating table syntax
					TABLE  (TABLEstmt)             for creating table syntax
					(      (OPENEDP)
                    )      (CLOSEDP)
					,      (COMMA)
					[      (OPENEDB)
					]      (CLOSEDB)
					INSERT (INSERTINGstmt)         for the inserting syntax
					INTO   (INTOstmt)              for the inserting syntax
					VALUES (VALUESstmt)            for the inserting syntax
					'      (APOSTROPHE)
					*      (STAR)                  refering to all or forevrything
					FROM   (FROMstmt)              queries syntax
					SELECT (SELECTstmt)            queries syntax
					WHERE  (WHEREstmt)             queries syntax
					>      (LESS_THAN)
					<      (GREATER_THAN)
					!=     (NOT_EQUAL)
					==     (EQUAL)
				    >=     (GREATER_THAN_OR_EQUAL)
					<=     (LESS_THAN_OR_EQUAL)

					\n     (LINE_END)  i dont  think it really matter cause the ; is the line end for our lang


)

IDENTIFIER (IDEN) : such as university or MyDataBase1 ..ect anynames ( must start with letter followed
                                                                by a letter or digit )
NUMBER      such as 12 55 44 1.5 12.5
I_string      have evrything letter digit underscore ..ect

*/
/*

PROGRAMME Universite;
CREATE TABLE Etudiant (
id INTEGER,
age INTEGER,
moyenne REAL
);
INSERT INTO Etudiant VALUES ();
INSERT INTO Etudiant VALUES ();
INSERT INTO Etudiant VALUES ();
SELECT * FROM Etudiant;
SELECT id, moyenne FROM Etudiant;
SELECT * FROM Etudiant WHERE age > 20;
FIN


S -> startup T final
startup ->  STARTstmt IDEN FINALstmt
final -> ENDstmt
T -> empty
     input;
input ->
      empty
	  line input;
line ->
       creatingstmtNT   NT MEANS NON TERMINAL
	   insertionstmtNT
	   querylinesNT
	   empty
creatingstmtNT ->
       TABLECREATINGstmt TABLEstmt IDEN OPENEDP decNT CLOSEDP FINALstmt
decNT ->
       IDENT typeNT
	   IDENT typeNT COMMA IDENT typeNT
typeNT ->
       INTEGER
	   REAL
	   string

insertionstmtNT ->
       INSERTstmt INTOstmt IDENT VALUESstmt OPENEDP insertDec CLOSEDP FINALstmt
insertDec ->
       inputsTypeNT
	   inputsTypeNT COMMA inputsTypeNT
querylinesNT ->
       SELECTstmt selectProbNT FROMstmt IDENT locationNT FINALstmt
selectProbNT ->
       STAR
	   identNT
identNT ->
       IDENT
	   IDENT COMMA IDENT
locationNT ->
       empty
       WHEREstmt whereThingsNT operatorNT whereThingsNT
whereThingsNT ->
       inputsTypeNT
	   IDENT
operatorNT ->
       LESS_THAN
	   GREATER_THAN
	   NOT_EQUAL
	   EQUAL
	   GREATER_THAN_OR_EQUAL
	   LESS_THAN_OR_EQUAL



inputsTypeNT ->
       NUMBER
	   I_string






*/


type Kind int



const (
	// --------- Keywords -------------

	UNKNOWN_TOKEN       Kind = iota
	STARTstmt                   // PROGRAMME
	FINALstmt                   // ;
	ENDstmt                     // FIN
	TABLECREATINGstmt           // CREATE
	TABLEstmt                   // TABLE
	INSERTINGstmt               // INSERT 
	INTOstmt                    // INTO
	VALUESstmt                  // VALUES
	SELECTstmt                  // SELECT
	FROMstmt                    // FROM
	WHEREstmt                   // WHERE
	TYPEINTstmt                 // INTEGER
	TYPEFLOATstmt               // REAL
	TYPESTRINGstmt              // string

	// ---------- Punctuation ----------

	OPENEDP                     // (
	CLOSEDP                     // )
	COMMA                       // ,
	OPENEDB                     // [
	CLOSEDB                     // ]

	// ------------ Operators ------------

	LESS_THAN                   // <
	GREATER_THAN                // >
	NOT_EQUAL                   // != 
	EQUAL                       // = 
	LESS_THAN_OR_EQUAL          // <= 
	GREATER_THAN_OR_EQUAL       // >=
	STAR                        // *

	// ------------ Literals -------------

	IDEN                        // for example : university or hello1 or MyDataBase1 ...ect
	NUMBER_INT                  // INTERGER NUMBERS for examples : 12 22 0 12 18 177 
	NUMBER_REAL                 // FLOATING-POINT NUMBERS for examples : 1.5 0.2 15.55 13.6666
	I_string                    // a whole input including the quotes 'benali' 'omar' 'example@gmail.com'
	EOF                         // END OF FILE 
	LEXICAL_ERROR               // anycaracter that isnt a token 
)

type Token struct {
	Literal string              // The Literal token value like ; FIN INSERT
	line int                    // line of the token
	column int                  // column of the token
	TokenType Kind              // the token type like EQUAL OR STAR OR ect
} 

var KeyWords = map[string]Kind {
	"PROGRAMME":        STARTstmt,                        
	"FIN":              ENDstmt,                    
	"CREATE":           TABLECREATINGstmt,          
	"TABLE":            TABLEstmt,                   
	"INSERT":           INSERTINGstmt,               
	"INTO":             INTOstmt,                    
	"VALUES":           VALUESstmt,                  
	"SELECT":           SELECTstmt,                  
	"FROM":             FROMstmt,                    
	"WHERE":            WHEREstmt,                   
	"INTEGER":          TYPEINTstmt,                 
	"REAL":             TYPEFLOATstmt,               
	"string":           TYPESTRINGstmt,	
}

func (t Token) TokenInfo() string {
	return fmt.Sprintf("Token{ Type : %v , Value : %v , line : %v, column : %v }", t.TokenType, t.Literal, t.line, t.column);
}



