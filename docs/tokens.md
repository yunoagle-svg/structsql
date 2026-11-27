# StructSQL — Token Reference

Final token list, using your own names. See `StructSQL_grammar_tokens.pdf`
section 1 for what changed from the draft and why.

> **Changelog (latest):** reversed an earlier call on this page —
> keywords are **case-sensitive**, not case-insensitive. `create` and
> `Create` are plain `IDEN`, only the exact spelling `CREATE` is
> `TABLECREATINGstmt` (confirmed against the actual lexer: `create` comes
> back as `IDEN`, which is correct, not a bug). Also dropped the curly-quote
> `’` as a valid string delimiter — only the straight `'` (U+0027) is
> supported, since the lexer indexes the source as raw bytes (no `rune`
> handling), and `’` is a 3-byte UTF-8 sequence that byte indexing can't
> safely treat as a single quote character.

## Lexer rules (apply before tokenizing)

- **Whitespace** (space, tab, `\n`, `\r`) is skipped, not tokenized. `;` is
  the real statement terminator, so newlines carry no meaning.
- **Keywords are case-sensitive.** Every keyword must match the exact
  spelling in the table below — `CREATE` is a keyword, `create`/`Create`/
  `CREATE ` with different casing are just an `IDEN`. No uppercasing or
  normalizing of the lexeme before the keyword lookup.
- **String quotes:** only the straight `'` (U+0027) delimits a string. The
  curly quote `’` is not supported — treat it as an unrecognized character
  if it shows up (lexical error).
- On an unrecognized character, report a lexical error and keep scanning —
  don't stop at the first bad character.

## Keywords (case-sensitive)

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

| Token           | Meaning                                                                                                                                        |
| --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `UNKNOWN_TOKEN` | the zero value of the `Kind` type — never emitted on purpose. If a token ever shows up with this kind, something didn't get filled in.         |
| `EOF`           | end of the source file — appended once, as the last token in the slice                                                                         |
| `LEXICAL_ERROR` | an unrecognized character. Reported (error printed) but **does not stop scanning** — the lexer keeps producing tokens for everything after it. |

`\n` is **not** a token — confirmed correct in the draft: whitespace is
skipped, and `;` is what actually ends a statement.
