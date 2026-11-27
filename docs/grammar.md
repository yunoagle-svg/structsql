# StructSQL — FIRST & FOLLOW Sets

LL(1) grammar verification for the StructSQL mini-compiler.

---

## Nonterminals (21)

```
S, STARTING, ENDING, stmts, line, CREATING, CREATINGEXP, COMMACASE_C,
TYPE, INSERTING, INSERTIONEXP, LITERALVALUES, COMMACASES_I, QUERIES,
SELECTEXP, COMMACASE_SE, FROMEXP, WHERECASE, CONDITON, CONDITIONEXP, OP
```

---

## FIRST Sets

| Nonterminal     | FIRST                                                                                      |
| --------------- | ------------------------------------------------------------------------------------------ |
| `S`             | `{ STARTstmt }`                                                                            |
| `STARTING`      | `{ STARTstmt }`                                                                            |
| `ENDING`        | `{ ENDstmt }`                                                                              |
| `stmts`         | `{ TABLECREATINGstmt, INSERTINGstmt, SELECTstmt, ε }`                                      |
| `line`          | `{ TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }`                                         |
| `CREATING`      | `{ TABLECREATINGstmt }`                                                                    |
| `CREATINGEXP`   | `{ IDEN }`                                                                                 |
| `COMMACASE_C`   | `{ COMMA, ε }`                                                                             |
| `TYPE`          | `{ TYPEINTstmt, TYPEFLOATstmt, TYPESTRINGstmt }`                                           |
| `INSERTING`     | `{ INSERTINGstmt }`                                                                        |
| `INSERTIONEXP`  | `{ NUMBER_INT, NUMBER_REAL, I_string }`                                                    |
| `LITERALVALUES` | `{ NUMBER_INT, NUMBER_REAL, I_string }`                                                    |
| `COMMACASES_I`  | `{ COMMA, ε }`                                                                             |
| `QUERIES`       | `{ SELECTstmt }`                                                                           |
| `SELECTEXP`     | `{ STAR, IDEN }`                                                                           |
| `COMMACASE_SE`  | `{ COMMA, ε }`                                                                             |
| `FROMEXP`       | `{ IDEN }`                                                                                 |
| `WHERECASE`     | `{ WHEREstmt, ε }`                                                                         |
| `CONDITON`      | `{ NUMBER_INT, NUMBER_REAL, I_string, IDEN }`                                              |
| `CONDITIONEXP`  | `{ NUMBER_INT, NUMBER_REAL, I_string, IDEN }`                                              |
| `OP`            | `{ LESS_THAN, GREATER_THAN, NOT_EQUAL, EQUAL, LESS_THAN_OR_EQUAL, GREATER_THAN_OR_EQUAL }` |

---

## FOLLOW Sets

| Nonterminal     | FOLLOW                                                                                                                                              |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `S`             | `{ EOF }`                                                                                                                                           |
| `STARTING`      | `{ TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }`                                                                                                  |
| `ENDING`        | `{ EOF }`                                                                                                                                           |
| `stmts`         | `{ ENDstmt }`                                                                                                                                       |
| `line`          | `{ TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }`                                                                                                  |
| `CREATING`      | `{ TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }`                                                                                                  |
| `CREATINGEXP`   | `{ CLOSEDP }`                                                                                                                                       |
| `COMMACASE_C`   | `{ CLOSEDP }`                                                                                                                                       |
| `TYPE`          | `{ COMMA }`                                                                                                                                         |
| `INSERTING`     | `{ TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }`                                                                                                  |
| `INSERTIONEXP`  | `{ CLOSEDP }`                                                                                                                                       |
| `LITERALVALUES` | `{ COMMA }`                                                                                                                                         |
| `COMMACASES_I`  | `{ CLOSEDP }`                                                                                                                                       |
| `QUERIES`       | `{ TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }`                                                                                                  |
| `SELECTEXP`     | `{ FROMstmt }`                                                                                                                                      |
| `COMMACASE_SE`  | `{ FROMstmt }`                                                                                                                                      |
| `FROMEXP`       | `{ WHEREstmt }`                                                                                                                                     |
| `WHERECASE`     | `{ FINALstmt, TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }`                                                                                       |
| `CONDITON`      | `{ FINALstmt, TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }`                                                                                       |
| `CONDITIONEXP`  | `{ LESS_THAN, GREATER_THAN, NOT_EQUAL, EQUAL, LESS_THAN_OR_EQUAL, GREATER_THAN_OR_EQUAL, FINALstmt, TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }` |
| `OP`            | `{ NUMBER_INT, NUMBER_REAL, I_string, IDEN }`                                                                                                       |

---

## LL(1) Verification

| Check                                   | Result                              |
| --------------------------------------- | ----------------------------------- |
| Left recursion                          | ✅ None (all lists right-recursive) |
| FIRST/FIRST conflicts                   | ✅ None                             |
| FIRST/FOLLOW conflicts on ε-productions | ✅ None                             |
| **Grammar is LL(1)**                    | ✅ **Yes**                          |

### ε-production safety

| Nonterminal    | ε-production           | FOLLOW                                                        | Conflict? |
| -------------- | ---------------------- | ------------------------------------------------------------- | --------- |
| `stmts`        | `stmts → EMPTY`        | `{ ENDstmt }`                                                 | ✅        |
| `COMMACASE_C`  | `COMMACASE_C → EMPTY`  | `{ CLOSEDP }`                                                 | ✅        |
| `COMMACASES_I` | `COMMACASES_I → EMPTY` | `{ CLOSEDP }`                                                 | ✅        |
| `COMMACASE_SE` | `COMMACASE_SE → EMPTY` | `{ FROMstmt }`                                                | ✅        |
| `WHERECASE`    | `WHERECASE → EMPTY`    | `{ FINALstmt, TABLECREATINGstmt, INSERTINGstmt, SELECTstmt }` | ✅        |
