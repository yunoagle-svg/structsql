# StructSQL — Token Reference

This is the final token list for the lexer. Every token the parser sees comes
from this table — nothing else.

## Lexer rules (apply before tokenizing)

- **Whitespace** (space, tab, `\n`, `\r`) is skipped, not tokenized. `;` is the
  real statement terminator, so newlines carry no meaning.
- **Keywords are case-insensitive.** `CREATE`, `create` and `Create` are the
  same token. Match the identifier/keyword scanner against an uppercased copy
  of the lexeme.
- **String quotes:** accept both `'` (U+0027) and `’` (U+2019, the curly
  quote used in the spec's own PDF) as string delimiters.
- On an unrecognized character, emit an `ERROR lexical` (see
  `output_contract.md`) and keep scanning — don't stop at the first bad
  character.

## Keywords

| Token       | Lexeme (case-insensitive) | Meaning                      |
| ----------- | ------------------------- | ---------------------------- |
| `PROGRAMME` | `PROGRAMME`               | starts the program           |
| `FIN`       | `FIN`                     | ends the program             |
| `CREATE`    | `CREATE`                  | start of a table declaration |
| `TABLE`     | `TABLE`                   | part of `CREATE TABLE`       |
| `INSERT`    | `INSERT`                  | start of an insertion        |
| `INTO`      | `INTO`                    | part of `INSERT INTO`        |
| `VALUES`    | `VALUES`                  | part of `INSERT ... VALUES`  |
| `SELECT`    | `SELECT`                  | start of a query             |
| `FROM`      | `FROM`                    | part of `SELECT ... FROM`    |
| `WHERE`     | `WHERE`                   | start of a condition         |

## Types

| Token        | Lexeme (case-insensitive) | Meaning                                                |
| ------------ | ------------------------- | ------------------------------------------------------ |
| `INTEGER`    | `INTEGER`                 | whole number type                                      |
| `REAL`       | `REAL`                    | floating-point type                                    |
| `STRINGTYPE` | `string`                  | character-string type (optionally sized: `string[20]`) |

## Punctuation / separators

| Token      | Lexeme | Meaning                          |
| ---------- | ------ | -------------------------------- |
| `SEMI`     | `;`    | statement terminator             |
| `LPAREN`   | `(`    | opens an attribute/value list    |
| `RPAREN`   | `)`    | closes an attribute/value list   |
| `COMMA`    | `,`    | separates list items             |
| `LBRACKET` | `[`    | opens a string size, e.g. `[20]` |
| `RBRACKET` | `]`    | closes a string size             |

## Operators

| Token  | Lexeme | Meaning                 |
| ------ | ------ | ----------------------- |
| `EQ`   | `=`    | equal                   |
| `NEQ`  | `!=`   | not equal               |
| `LT`   | `<`    | less than               |
| `GT`   | `>`    | greater than            |
| `LE`   | `<=`   | less than or equal      |
| `GE`   | `>=`   | greater than or equal   |
| `STAR` | `*`    | "all columns" in SELECT |

> Double-check these against your scanner: `<=` and `>=` must be read as one
> two-character token (the DFA peeks at the next character before deciding
> between `<` and `<=`, same for `>`/`>=`).

## Literals

| Token      | Example lexemes        | Rule                                                                                                                                                                                                 |
| ---------- | ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `IDENT`    | `Etudiant`, `moyenne1` | starts with a letter, then letters or digits                                                                                                                                                         |
| `NUM_INT`  | `12`, `55`, `0`        | digits only, no `.`                                                                                                                                                                                  |
| `NUM_REAL` | `1.5`, `14.5`, `20.0`  | digits, one `.`, then digits                                                                                                                                                                         |
| `STRING`   | `'benali'`             | opening quote, any characters except the quote, closing quote. The quotes themselves are **not** separate tokens — the whole thing is one `STRING` token whose value is the text between the quotes. |

`NUM_INT` and `NUM_REAL` are two different token kinds, not one `NUMBER`
kind. Splitting them here means the semantic analyzer doesn't have to
re-inspect the lexeme text later to know whether a value is an integer or a
real.

## Special

| Token   | Meaning                                                                        |
| ------- | ------------------------------------------------------------------------------ |
| `EOF`   | end of the source file                                                         |
| `ERROR` | an unrecognized character (see `output_contract.md` for the error line format) |

Claude finished the response
