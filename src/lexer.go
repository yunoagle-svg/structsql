// Hand-written lexer (DFA coded by hand): reads the source character by character and returns tokens; reports lexical errors with line and column.
package main

import (
	"fmt"
	"strconv"
)

type lexer struct {
	source   string  // the source code in string
	position int     // position of the pointer when we move
	line     int     // line
	column   int     // column
}

func (t *Token) fillup(lit string, line int, column int, kind Kind) {
	/*

	just infos fillup for the token struct

	*/
	t.Literal = lit
	t.line = line
	t.column = column
	t.TokenType = kind
}

func isDigit(s string) bool {
	/*check if a digit specially integer */
	_, ok := strconv.Atoi(s)
	if ok != nil {
		return false
	}
	return true
}

func isFloat(s string) bool {
	/* check if a floating-point number */
	_, ok := strconv.ParseFloat(s, 64)
	if ok != nil {
		return false
	}
	return true
}

func isLetter(s string) bool {
	/* if its letter a-z and A-Z*/
	return len(s) == 1 &&
		((s[0] >= 'a' && s[0] <= 'z') ||
			(s[0] >= 'A' && s[0] <= 'Z'))
}

func createLexer(str string) *lexer {
	/* inisialize the constructor  */
	return &lexer{

		source:   str,
		position: 0,
		line:     0,
		column:   0,
	}
}

func isWhiteSpace(str string) (int, int, bool) {
	lcount, Ccounter := 0, 0
	if str == " " || str == "\t" {
		/*i add here the columns bc we are skipping the spaces and tabs so we moving in columns*/
		Ccounter++
		return lcount, Ccounter, true
	}
	if str == "\n" {
		/* i added this bc we skiping the endlines so we moving in lines*/
		lcount++
		return lcount, Ccounter, true
	}
	return lcount, Ccounter, false
}

func isEmpty(s string) bool {
	return s == ""
}

/*

PROGRAMME Universite;

*/
func (l *lexer) splitIntoTokens() ([]Token, error) {
	result := ""
	var tokens []Token

	for ; l.position < len(l.source); l.position++ {
		chStr := string(l.source[l.position])

		if lcounter, ccounter, ok := isWhiteSpace(chStr); ok { // wee skipp the white spaces and check for the result the word before the the space
			if !isEmpty(result) { // we check when if result isnt ampty so we check for tokens
				tok, err := check(result, l.line, l.column)
				
				//     if err != nil { return tokens, err }
				//     tokens = append(tokens, tok)
				// which stopped the WHOLE scan the first time any bad word
				// showed up. Now we keep the LEXICAL_ERROR token check()
				// already built, print the error, and keep going. Example:
				// for "create TABLE ... @@@ FIN", before this change FIN
				// never showed up in the output after the @@@ error — now
				// it will.
				tokens = append(tokens, tok)
				if err != nil {
					fmt.Println(err)
				}
				result = ""
			}
			l.line += lcounter // adding to the lexer structure the lines we skipped
			if lcounter > 0 {
				l.column = 0
			} else {
				l.column += ccounter // adding to the lexer structure the columns we skipped
			}
			continue
		}

		/*
		here we have two cases the operators and symbols in len of 1 and len of 2
		first len = 2 case
		so we check before if the current position + 1 isnt >= the len(source)
		then we take the the two symbols and store them in a temp and check for the word or token
		before the operatore and add it than check for the temp  that hold the operator
		for example we have

		                   | a | g | e | < | = | 1 | 8
						   or
						   | U | n | i | v | e | r | s | i | t | e | ; EOS (END OF STRING)
		*/
		if l.position+1 < len(l.source) { // here it check for the word before the symbol token case of len -> 2
			twoChar := l.source[l.position : l.position+2]
			if symKind, ok := isSymbol(twoChar); ok {
				if !isEmpty(result) {
					tok, err := check(result, l.line, l.column)
					// CHANGED: same reasoning as above — keep scanning
					// instead of returning out of the whole function.
					tokens = append(tokens, tok)
					if err != nil {
						fmt.Println(err)
					}
					result = ""
				}

				var symTok Token
				symTok.fillup(twoChar, l.line, l.column, symKind)
				tokens = append(tokens, symTok)
				l.position++ // consume second character of the 2-char symbol
				l.column += 2
				continue
			}
		}

		if symKind, ok := isSymbol(chStr); ok {
			/* here if its a symbol of len -> 1 */
			if !isEmpty(result) {
				tok, err := check(result, l.line, l.column)
				// CHANGED: same reasoning as above.
				tokens = append(tokens, tok)
				if err != nil {
					fmt.Println(err)
				}
				result = ""
			}

			var symTok Token
			symTok.fillup(chStr, l.line, l.column, symKind)
			tokens = append(tokens, symTok)
			l.column++
			continue
		}

		result += chStr
		l.column++
	}

	if !isEmpty(result) {
		tok, err := check(result, l.line, l.column)
		// same reasoning — this is the last word in the file (no
		// trailing whitespace/symbol to trigger a flush), so it needs the
		// same "keep going, don't bail" treatment as the other three spots.
		tokens = append(tokens, tok)
		if err != nil {
			fmt.Println(err)
		}
	}

	//  One EOF token appended at the very
	//  end so the parser (or anything printing the token list) can see where
	//  the stream actually ends, e.g. for "FIN" you now get
	//  ...{Literal:FIN ... TokenType:ENDstmt}, {Literal:"" ... TokenType:EOF}
	//  as the last entry.
	tokens = append(tokens, Token{Literal: "", line: l.line, column: l.column, TokenType: EOF})

	return tokens, nil
	//always nil now, since every lexical error is already reported
	// (fmt.Println above) and recorded in-stream as a LEXICAL_ERROR token at
	// the exact spot it happened, instead of aborting the whole function
	// with one error value.
}

func check(str string, line int, column int) (Token, error) {
	/*
		Graping each word see if its a valid token or no

	*/
	var tok Token
	if val, ok := KeyWords[str]; ok {
		/*First Check if its an Keyword */
		tok.fillup(str, line, column, val)
		return tok, nil
	} else if isIdent(str) {
		/*Then if an ident*/
		tok.fillup(str, line, column, IDEN)
		return tok, nil
	} else if isDigit(str) {
		/* if its integer */
		tok.fillup(str, line, column, NUMBER_INT)
		return tok, nil
	} else if isFloat(str) {
		/* if its floating-point number */
		tok.fillup(str, line, column, NUMBER_REAL)
		return tok, nil
	} else if isI_Sting(str) {
		/*if its an iput string if its like 'binali' or 'omar' ..ect*/
		tok.fillup(str, line, column, I_string)
		return tok, nil
	} else if val, ok := isSymbol(str); ok {
		/* if its symbol like < > = != * ; ( ) */
		tok.fillup(str, line, column, val)
		return tok, nil
	} else {
		/* Unreconized token error */
		tok.fillup(str, line, column, LEXICAL_ERROR)
		return tok, fmt.Errorf("lexical error (Unknown token -> ' %v ') in line %v column %v ", str, line, column)
	}
}

func isI_Sting(s string) bool {
	/*
	check for the I_string such as  'kdjsqhd' 'omar' 'example@gmail.com' like
	*/
	if len(s) < 2 || s[0] != '\'' || s[len(s)-1] != '\'' {
		return false
	}
	return true
}

func isSymbol(str string) (Kind, bool) {
	/* check if ites a symbol */
	if len(str) == 0 || len(str) >= 3 {
		return LEXICAL_ERROR, false
	}

	switch str {
	case "(":
		return OPENEDP, true
	case ")":
		return CLOSEDP, true
	case ",":
		return COMMA, true
	case "[":
		return OPENEDB, true
	case "]":
		return CLOSEDB, true
	case "<":
		return LESS_THAN, true
	case ">":
		return GREATER_THAN, true
	case "!=":
		return NOT_EQUAL, true
	case "=":
		return EQUAL, true
	case "<=":
		return LESS_THAN_OR_EQUAL, true
	case ">=":
		return GREATER_THAN_OR_EQUAL, true
	case "*":
		return STAR, true
	case ";":
		return FINALstmt, true
	default:
		return LEXICAL_ERROR, false
	}
}

func isIdent(str string) bool {
	/* check if identifier and accept only start with a letter and followed by letter or digit
	no symbols */

	if len(str) == 0 || !isLetter(string(str[0])) {
		return false
	}

	for _, ch := range str {
		s := string(ch)
		if !isDigit(s) && !isLetter(s) {
			return false
		}
	}
	return true
}
func main() {
		s := `create TABLE Etudiant (id INTEGER, nom string[20]); @@@ FIN`
		l := createLexer(s)
		tokens, err := l.splitIntoTokens()
		if err != nil {
			fmt.Println("Error:", err)
		}
		for _, token := range tokens {
			fmt.Printf("%+v\n", token)
		}
}