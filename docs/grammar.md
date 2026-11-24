# StructSQL — Grammar

LL(1), predictive, table-driven. Token names match `tokens.md`.

## 1. Design decisions baked into this grammar

- **Declarations before operations.** The spec requires every `CREATE TABLE`
  to come before any `INSERT`/`SELECT`. Instead of one repeating "statement"
  rule that would accept them in any order, the grammar has two separate
  blocks: `Decls` (only `CreateStmt`s) followed by `Ops` (only
  `InsertStmt`/`SelectStmt`). Mixing the order is then a **syntax** error,
  not something semantic analysis has to catch.
- **Comma lists are recursive**, not capped at two items. An attribute list,
  a VALUES list and a SELECT column list can all have any number of items —
  `AttrRest`, `ValRest` and `IdentRest` each recurse on `COMMA`.
- **WHERE is deliberately narrow:** `IDENT RelOp Value` only — an attribute
  compared to a literal. The spec's example (`WHERE age > 20`) never compares
  two attributes or two literals, so the grammar doesn't allow it either.
- **`string[20]` is optional.** `SizeOpt` is `ε` when no size is given.

## 2. Grammar

```
Program     -> PROGRAMME IDENT SEMI Decls Ops FIN

Decls       -> CreateStmt Decls
             | ε

CreateStmt  -> CREATE TABLE IDENT LPAREN AttrList RPAREN SEMI

AttrList    -> IDENT Type AttrRest
AttrRest    -> COMMA IDENT Type AttrRest
             | ε

Type        -> INTEGER
             | REAL
             | STRINGTYPE SizeOpt
SizeOpt     -> LBRACKET NUM_INT RBRACKET
             | ε

Ops         -> OpStmt Ops
             | ε
OpStmt      -> InsertStmt
             | SelectStmt

InsertStmt  -> INSERT INTO IDENT VALUES LPAREN ValList RPAREN SEMI
ValList     -> Value ValRest
ValRest     -> COMMA Value ValRest
             | ε
Value       -> NUM_INT
             | NUM_REAL
             | STRING

SelectStmt  -> SELECT SelectList FROM IDENT WhereOpt SEMI
SelectList  -> STAR
             | IdentList
IdentList   -> IDENT IdentRest
IdentRest   -> COMMA IDENT IdentRest
             | ε
WhereOpt    -> WHERE IDENT RelOp Value
             | ε
RelOp       -> EQ | NEQ | LT | GT | LE | GE
```

## 3. FIRST sets

| Non-terminal | FIRST                               |
| ------------ | ----------------------------------- |
| Program      | `PROGRAMME`                         |
| Decls        | `CREATE`, ε                         |
| CreateStmt   | `CREATE`                            |
| AttrList     | `IDENT`                             |
| AttrRest     | `COMMA`, ε                          |
| Type         | `INTEGER`, `REAL`, `STRINGTYPE`     |
| SizeOpt      | `LBRACKET`, ε                       |
| Ops          | `INSERT`, `SELECT`, ε               |
| OpStmt       | `INSERT`, `SELECT`                  |
| InsertStmt   | `INSERT`                            |
| ValList      | `NUM_INT`, `NUM_REAL`, `STRING`     |
| ValRest      | `COMMA`, ε                          |
| Value        | `NUM_INT`, `NUM_REAL`, `STRING`     |
| SelectStmt   | `SELECT`                            |
| SelectList   | `STAR`, `IDENT`                     |
| IdentList    | `IDENT`                             |
| IdentRest    | `COMMA`, ε                          |
| WhereOpt     | `WHERE`, ε                          |
| RelOp        | `EQ`, `NEQ`, `LT`, `GT`, `LE`, `GE` |

## 4. FOLLOW sets

| Non-terminal | FOLLOW                              |
| ------------ | ----------------------------------- |
| Program      | `$` (end of input)                  |
| Decls        | `INSERT`, `SELECT`, `FIN`           |
| CreateStmt   | `CREATE`, `INSERT`, `SELECT`, `FIN` |
| AttrList     | `RPAREN`                            |
| AttrRest     | `RPAREN`                            |
| Type         | `COMMA`, `RPAREN`                   |
| SizeOpt      | `COMMA`, `RPAREN`                   |
| Ops          | `FIN`                               |
| OpStmt       | `INSERT`, `SELECT`, `FIN`           |
| InsertStmt   | `INSERT`, `SELECT`, `FIN`           |
| ValList      | `RPAREN`                            |
| ValRest      | `RPAREN`                            |
| Value        | `COMMA`, `RPAREN`, `SEMI`           |
| SelectStmt   | `INSERT`, `SELECT`, `FIN`           |
| SelectList   | `FROM`                              |
| IdentList    | `FROM`                              |
| IdentRest    | `FROM`                              |
| WhereOpt     | `SEMI`                              |
| RelOp        | `NUM_INT`, `NUM_REAL`, `STRING`     |

## 5. Why this is LL(1)

For every non-terminal with more than one production, the FIRST sets of its
alternatives are disjoint — check each row of section 3: no token appears in
two alternatives of the same non-terminal (and where one alternative is ε,
its FOLLOW set in section 4 never overlaps the other alternatives' FIRST
set either). That's the condition for a table-driven predictive parser:
looking at one lookahead token always tells you which production to use,
with no backtracking.

## 6. Building the parse table

The table is `Table[NonTerminal][Token] -> Production`. Fill it by this rule,
for every production `A -> α`:

- For every terminal `t` in `FIRST(α)`, set `Table[A][t] = A -> α`.
- If `ε` is in `FIRST(α)`, also set `Table[A][t] = A -> α` for every `t` in
  `FOLLOW(A)`.
  Because of section 5, no cell ever gets written twice. This table is what
  `parse_table.go` encodes as data — build it once from the sets above, by
  hand, and keep it as a literal table (a map or 2D array), not re-derived at
  runtime.

## 7. Worked trace (sanity check)

```
PROGRAMME Universite;
CREATE TABLE Etudiant (id INTEGER, nom STRINGTYPE[20], age INTEGER, moyenne REAL);
INSERT INTO Etudiant VALUES (1, 'benali', 21, 14.5);
SELECT * FROM Etudiant;
SELECT id, moyenne FROM Etudiant WHERE age > 20;
FIN
```

- `Program` consumes `PROGRAMME IDENT SEMI`, then enters `Decls`.
- Lookahead `CREATE` → `Decls -> CreateStmt Decls`. The `CreateStmt` consumes
  the whole table declaration; `AttrRest` recurses three times (`nom`, `age`,
  `moyenne`) before hitting `)` and reducing to ε.
- Lookahead after that `CreateStmt` is `INSERT`, which is not in
  `FIRST(CreateStmt)` (`CREATE` only) — so `Decls -> ε` is chosen, and we
  drop into `Ops`.
- `Ops` sees `INSERT` → `OpStmt -> InsertStmt`. `ValRest` recurses three
  times for the four values, then hits `)`.
- Next `Ops` sees `SELECT` → `OpStmt -> SelectStmt` (twice, one with `WHERE`,
  one without — `WhereOpt -> ε` the first time, `WhereOpt -> WHERE IDENT
RelOp Value` the second).
- Lookahead `FIN` is not in `FIRST(OpStmt)` → `Ops -> ε`, then `Program`
  consumes `FIN`. Done.
  If you can walk through this by hand with the parse table from section 6,
  your table is built correctly.

Claude finished the response
