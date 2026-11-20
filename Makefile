# StructSQL mini-compiler - build script
#
#   make          build the compiler  -> bin/structsql
#   make test     run the compiler on every file in tests/ (valid.sql must pass, err_*.sql must fail)
#   make demo     compile tests/valid.sql, build the generated C program and run it (writes data/*.dat)
#   make clean    remove build files, generated program and data files (keeps the .gitkeep files)

CC      = gcc
CFLAGS  = -std=c11 -Wall -Wextra -g

SRC     = $(wildcard src/*.c)
HDR     = $(wildcard src/*.h)
OBJ     = $(patsubst src/%.c,bin/obj/%.o,$(SRC))
BIN     = bin/structsql

.PHONY: all test demo clean

all: $(BIN)

$(BIN): $(OBJ)
	$(CC) $(CFLAGS) -o $@ $^

# every object depends on all headers (small project, simplest safe rule)
bin/obj/%.o: src/%.c $(HDR)
	@mkdir -p bin/obj
	$(CC) $(CFLAGS) -c $< -o $@

test: all
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
	$(CC) $(CFLAGS) -o generated/out generated/out.c
	cd data && ../generated/out

clean:
	rm -rf bin/obj $(BIN)
	rm -f generated/out generated/out.c data/*.dat
