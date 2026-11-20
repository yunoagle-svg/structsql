/* Lexer interface: token types enum, Token struct, lexer init and next_token() prototypes. */


enum Tokens { /* all token we have in the language */
    PROGRAM,       // "PROGRAM" start of the program 
    IDENT,         // identifier that u gonna name by it variables tables ...ect
    CREATE,        // "CREATE" keyword for creating tables
    TABLE,         // keyword 
    INTEGER,       // keyword
    REAL,          // keyword
    INSERT,        // keyword
    INTO,          // keyword
    VALUES,        // keyword
    OPENEDP,       // (
    CLOSEDP,       // )
    VERG,          // ,
    SELECT,        // keyword
    FROM,          // keyword
    ALL,           // *
    FINAL,         // ;
    EOL,           // \n
    COMPARATOR,    // == < > >= <= != 
    FIN            // keyword 
};


