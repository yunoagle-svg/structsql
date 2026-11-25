# StructSQL — Grammar

LL(1), predictive, table-driven, using your own non-terminal names. See
`StructSQL_grammar_tokens.pdf` section 1 for what changed from the draft
and why — short version:

- `declsNT` (only table creation, repeating) now comes strictly before
  `opsNT` (only insert/select, repeating), enforcing "declarations before
  operations." The old `T`/`input`/`line` also added a stray extra `;` on
  top of each statement's own terminator — gone now.
- `decRestNT`, `insertDecRest`, `identRestNT` are new: they recurse on
  `COMMA` so attribute lists, VALUES lists and SELECT column lists can have
  **any** number of items, not just 1 or 2.
- `whereThingsNT` is gone. `locationNT` now requires `IDEN operatorNT
inputsTypeNT` directly — left side is always an attribute, right side is
  always a literal.
- `sizeNT` is new, so `typeNT` actually uses `OPENEDB`/`CLOSEDB` for
  `string[20]`.

## 1. Grammar

```
S         -> startup T final
startup   -> STARTstmt IDEN FINALstmt
final     -> ENDstmt

T         -> declsNT opsNT

declsNT   -> creatingstmtNT declsNT
           | empty
opsNT     -> lineNT opsNT
           | empty
lineNT    -> insertionstmtNT
           | querylinesNT

creatingstmtNT -> TABLECREATINGstmt TABLEstmt IDEN OPENEDP decNT CLOSEDP FINALstmt

decNT       -> IDEN typeNT decRestNT
decRestNT   -> COMMA IDEN typeNT decRestNT
             | empty
typeNT      -> TYPEINTstmt
             | TYPEFLOATstmt
             | TYPESTRINGstmt sizeNT
sizeNT      -> OPENEDB NUMBER_INT CLOSEDB
             | empty

insertionstmtNT -> INSERTINGstmt INTOstmt IDEN VALUESstmt OPENEDP insertDec CLOSEDP FINALstmt

insertDec      -> inputsTypeNT insertDecRest
insertDecRest  -> COMMA inputsTypeNT insertDecRest
               | empty
inputsTypeNT   -> NUMBER_INT
               | NUMBER_REAL
               | I_string

querylinesNT -> SELECTstmt selectProbNT FROMstmt IDEN locationNT FINALstmt

selectProbNT  -> STAR
              | identNT
identNT       -> IDEN identRestNT
identRestNT   -> COMMA IDEN identRestNT
              | empty

locationNT -> empty
            | WHEREstmt IDEN operatorNT inputsTypeNT

operatorNT -> LESS_THAN | GREATER_THAN | NOT_EQUAL
            | EQUAL | GREATER_THAN_OR_EQUAL | LESS_THAN_OR_EQUAL
```

## 2. FIRST sets

| Non-terminal    | FIRST                                                                                            |
| --------------- | ------------------------------------------------------------------------------------------------ |
| S               | `STARTstmt`                                                                                      |
| startup         | `STARTstmt`                                                                                      |
| final           | `ENDstmt`                                                                                        |
| T               | `TABLECREATINGstmt`, `INSERTINGstmt`, `SELECTstmt`, ε                                            |
| declsNT         | `TABLECREATINGstmt`, ε                                                                           |
| opsNT           | `INSERTINGstmt`, `SELECTstmt`, ε                                                                 |
| lineNT          | `INSERTINGstmt`, `SELECTstmt`                                                                    |
| creatingstmtNT  | `TABLECREATINGstmt`                                                                              |
| decNT           | `IDEN`                                                                                           |
| decRestNT       | `COMMA`, ε                                                                                       |
| typeNT          | `TYPEINTstmt`, `TYPEFLOATstmt`, `TYPESTRINGstmt`                                                 |
| sizeNT          | `OPENEDB`, ε                                                                                     |
| insertionstmtNT | `INSERTINGstmt`                                                                                  |
| insertDec       | `NUMBER_INT`, `NUMBER_REAL`, `I_string`                                                          |
| insertDecRest   | `COMMA`, ε                                                                                       |
| inputsTypeNT    | `NUMBER_INT`, `NUMBER_REAL`, `I_string`                                                          |
| querylinesNT    | `SELECTstmt`                                                                                     |
| selectProbNT    | `STAR`, `IDEN`                                                                                   |
| identNT         | `IDEN`                                                                                           |
| identRestNT     | `COMMA`, ε                                                                                       |
| locationNT      | `WHEREstmt`, ε                                                                                   |
| operatorNT      | `LESS_THAN`, `GREATER_THAN`, `NOT_EQUAL`, `EQUAL`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN_OR_EQUAL` |

## 3. FOLLOW sets

| Non-terminal    | FOLLOW                                                        |
| --------------- | ------------------------------------------------------------- |
| S               | `$` (end of input)                                            |
| startup         | `TABLECREATINGstmt`, `INSERTINGstmt`, `SELECTstmt`, `ENDstmt` |
| T               | `ENDstmt`                                                     |
| final           | `$`                                                           |
| declsNT         | `INSERTINGstmt`, `SELECTstmt`, `ENDstmt`                      |
| opsNT           | `ENDstmt`                                                     |
| creatingstmtNT  | `TABLECREATINGstmt`, `INSERTINGstmt`, `SELECTstmt`, `ENDstmt` |
| lineNT          | `INSERTINGstmt`, `SELECTstmt`, `ENDstmt`                      |
| insertionstmtNT | `INSERTINGstmt`, `SELECTstmt`, `ENDstmt`                      |
| querylinesNT    | `INSERTINGstmt`, `SELECTstmt`, `ENDstmt`                      |
| decNT           | `CLOSEDP`                                                     |
| decRestNT       | `CLOSEDP`                                                     |
| typeNT          | `COMMA`, `CLOSEDP`                                            |
| sizeNT          | `COMMA`, `CLOSEDP`                                            |
| insertDec       | `CLOSEDP`                                                     |
| insertDecRest   | `CLOSEDP`                                                     |
| inputsTypeNT    | `COMMA`, `CLOSEDP`, `FINALstmt`                               |
| selectProbNT    | `FROMstmt`                                                    |
| identNT         | `FROMstmt`                                                    |
| identRestNT     | `FROMstmt`                                                    |
| locationNT      | `FINALstmt`                                                   |
| operatorNT      | `NUMBER_INT`, `NUMBER_REAL`, `I_string`                       |

## 4. Why this is LL(1)

For every non-terminal with more than one alternative, the FIRST sets above
are pairwise disjoint — e.g. `declsNT`'s two alternatives start with
`TABLECREATINGstmt` vs. {ε, whose governing FOLLOW is `INSERTINGstmt`/
`SELECTstmt`/`ENDstmt`}, no overlap; `typeNT`'s three alternatives start
with three different tokens; `operatorNT`'s six alternatives are six
distinct single tokens. One token of lookahead always picks the right
production, no backtracking needed.

Build the parse table as `Table[NonTerminal][Token] -> Production`: for
production `A -> α`, write it into column `t` for every `t` in `FIRST(α)`;
if `α` can be empty, also write it into column `t` for every `t` in
`FOLLOW(A)`. Because of the disjointness above, no cell is ever written
twice.

## 5. Worked trace (sanity check)

```
PROGRAMME Universite;
CREATE TABLE Etudiant (id INTEGER, nom TYPESTRINGstmt[20], age INTEGER, moyenne REAL);
INSERT INTO Etudiant VALUES (1, 'benali', 21, 14.5);
SELECT * FROM Etudiant;
SELECT id, moyenne FROM Etudiant WHERE age > 20;
FIN
```

`S` consumes `startup` (`STARTstmt IDEN FINALstmt`), then enters `T`.
Lookahead `TABLECREATINGstmt` selects `declsNT -> creatingstmtNT declsNT`;
`decRestNT` recurses twice more for `age` and `moyenne` before `CLOSEDP`
ends it. The next lookahead is `INSERTINGstmt`, which is not in
`FIRST(creatingstmtNT)`, so `declsNT -> empty` and we fall into `opsNT`.
`opsNT` sees `INSERTINGstmt` → `insertionstmtNT` (`insertDecRest` recurses
three times for the four values), then `SELECTstmt` twice (`locationNT ->
empty` the first time, `locationNT -> WHEREstmt IDEN operatorNT
inputsTypeNT` the second). Lookahead `ENDstmt` is not in `FIRST(lineNT)`,
so `opsNT -> empty`, `T` ends, and `final` consumes `ENDstmt`. Done.

If you can walk through this by hand against your parse table, the table
is built correctly.
