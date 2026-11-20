# Test files

valid.sql: complete correct program (CREATE, INSERT, SELECT *, SELECT attrs, SELECT WHERE).
err_lex.sql: contains an invalid character or unterminated string (lexical error).
err_syntax.sql: missing ';' or malformed statement (syntax error).
err_sem_table.sql: uses a table that was not declared.
err_sem_attr.sql: uses an attribute that does not exist.
err_sem_type.sql: inserts a value of the wrong type.
err_sem_count.sql: wrong number of values in INSERT.
err_sem_where.sql: invalid WHERE condition (type mismatch).
