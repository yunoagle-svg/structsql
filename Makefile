# StructSQL mini-compiler - build script (Go compiler core)
#
#   make          build the compiler  -> bin/structsql
#   make test     go unit tests, then run the compiler on tests/*.sql (valid.sql must pass, err_*.sql must fail)
#   make demo     compile tests/valid.sql to C, build the generated program with gcc and run it (writes data/*.dat)
#   make fmt      format all Go code
#   make clean    remove the binary, generated program and data files (keeps the .gitkeep files)

GO  = go
CC  = gcc
BIN = bin/structsql

.PHONY: all test demo fmt clean

all: $(BIN)

$(BIN): $(wildcard src/*.go) go.mod
	@mkdir -p bin
	$(GO) build -o $(BIN) ./src

test: all
	$(GO) vet ./src
	$(GO) test ./src
	@fail=0; \
	./$(BIN) sem tests/valid.sql >/dev/null 2>&1 \
	    && echo "ok   tests/valid.sql" \
	    || { echo "FAIL tests/valid.sql should pass"; fail=1; }; \
	for f in tests/err_*.sql; do \
	    if ./$(BIN) sem $$f >/dev/null 2>&1; then \
	        echo "FAIL $$f should report errors"; fail=1; \
	    else \
	        echo "ok   $$f"; \
	    fi; \
	done; \
	[ $$fail -eq 0 ]

demo: all
	./$(BIN) gen tests/valid.sql generated/out.c
	$(CC) -Wall -Wextra -o generated/out generated/out.c
	cd data && ../generated/out

fmt:
	$(GO) fmt ./src

clean:
	rm -f $(BIN) generated/out generated/out.c data/*.dat
