# StructSQL — Token Reference

Final token list, using your own names. See `StructSQL_grammar_tokens.pdf`
section 1 for what changed from the draft and why.

## Lexer rules (apply before tokenizing)

- **Whitespace** (space, tab, `\n`, `\r`) is skipped, not tokenized. `;` is
  the real statement terminator, so newlines carry no meaning.
- **Keywords are case-insensitive.** `CREATE`, `create` and `Create` are the
  same token. Match the identifier/keyword scanner against an uppercased
  copy of the lexeme.
- **String quotes:** accept both `'` (U+0027) and `’` (U+2019, the curly
  quote used in the spec's own PDF) as string delimiters.
- On an unrecognized character, report a lexical error and keep scanning —
  don't stop at the first bad character.

## Keywords (case-insensitive)

| Token               | Lexeme      | Meaning                                                |
| ------------------- | ----------- | ------------------------------------------------------ |
| `STARTstmt`         | `PROGRAMME` | starts the program                                     |
| `FINALstmt`         | `;`         | statement terminator                                   |
| `ENDstmt`           | `FIN`       | ends the program                                       |
| `TABLECREATINGstmt` | `CREATE`    | starts a table declaration                             |
| `TABLEstmt`         | `TABLE`     | part of `CREATE TABLE`                                 |
| `INSERTINGstmt`     | `INSERT`    | starts an insertion                                    |
| `INTOstmt`          | `INTO`      | part of `INSERT INTO`                                  |
| `VALUESstmt`        | `VALUES`    | part of `INSERT ... VALUES`                            |
| `SELECTstmt`        | `SELECT`    | starts a query                                         |
| `FROMstmt`          | `FROM`      | part of `SELECT ... FROM`                              |
| `WHEREstmt`         | `WHERE`     | starts a condition                                     |
| `TYPEINTstmt`       | `INTEGER`   | whole-number type                                      |
| `TYPEFLOATstmt`     | `REAL`      | floating-point type                                    |
| `TYPESTRINGstmt`    | `string`    | character-string type (optionally sized: `string[20]`) |

## Punctuation

| Token     | Lexeme | Meaning                          |
| --------- | ------ | -------------------------------- |
| `OPENEDP` | `(`    | opens an attribute/value list    |
| `CLOSEDP` | `)`    | closes an attribute/value list   |
| `COMMA`   | `,`    | separates list items             |
| `OPENEDB` | `[`    | opens a string size, e.g. `[20]` |
| `CLOSEDB` | `]`    | closes a string size             |

## Operators

| Token                   | Lexeme | Meaning                 |
| ----------------------- | ------ | ----------------------- |
| `LESS_THAN`             | `<`    | less than               |
| `GREATER_THAN`          | `>`    | greater than            |
| `NOT_EQUAL`             | `!=`   | not equal               |
| `EQUAL`                 | `=`    | equal                   |
| `LESS_THAN_OR_EQUAL`    | `<=`   | less than or equal      |
| `GREATER_THAN_OR_EQUAL` | `>=`   | greater than or equal   |
| `STAR`                  | `*`    | "all columns" in SELECT |

> **Fixed from the draft:** `<` is `LESS_THAN` and `>` is `GREATER_THAN`
> (they were swapped). `EQUAL` is a single `=`, not `==`, to match the
> spec's allowed operator list.
>
> `<=` and `>=` are each read as one two-character token: the DFA peeks at
> the next character before deciding between `LESS_THAN`/`LESS_THAN_OR_EQUAL`
> (same for `GREATER_THAN`).

## Literals

| Token         | Example                | Rule                                                                                             |
| ------------- | ---------------------- | ------------------------------------------------------------------------------------------------ |
| `IDEN`        | `Etudiant`, `moyenne1` | starts with a letter, then letters or digits                                                     |
| `NUMBER_INT`  | `12`, `55`, `0`        | digits only, no `.`                                                                              |
| `NUMBER_REAL` | `1.5`, `14.5`, `20.0`  | digits, one `.`, then digits                                                                     |
| `I_string`    | `'benali'`             | opening quote, any characters except the quote, closing quote — the whole thing is **one** token |

> **Fixed from the draft:** `NUMBER` is split into `NUMBER_INT` and
> `NUMBER_REAL` — decided once at lex time (by whether a `.` appears),
> instead of making semantic analysis re-inspect the lexeme text later.
>
> `APOSTROPHE` is **not** a separate token anymore. It was in the draft's
> token list but never used anywhere in the grammar, which meant strings
> couldn't actually be parsed. The lexer's DFA still watches for the quote
> character to know where a string starts and ends, it just doesn't emit
> the quote itself as a token — the whole `'...'` becomes one `I_string`.

## Special

| Token         | Meaning                                                                        |
| ------------- | ------------------------------------------------------------------------------ |
| `EOF`         | end of the source file                                                         |
| lexical error | an unrecognized character (see `output_contract.md` for the error line format) |

`\n` is **not** a token — confirmed correct in the draft: whitespace is
skipped, and `;` is what actually ends a statement.
