## personal dev is on macOS, but this should be
## supporting Linux as well
## gfind can be installed with `brew install findutils`
UNAME := $(shell uname)
ifeq ($(UNAME),Linux)
	FIND_BIN=find
endif
ifeq ($(UNAME),Darwin)
	FIND_BIN=gfind
endif

.SHELLFLAGS = -ec
.ONESHELL:

build: ./bin/pokeapi-cli

./bin/pokeapi-cli: $(shell $(FIND_BIN) -name "*.go")
	go build -o ./bin/pokeapi-cli

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run
run:
	PA_POKEAPI_POKEDEX_URL="https://pokeapi.co/api/v2/pokedex" \
	./bin/pokeapi-cli

./bin/gotestsum:
	GOBIN=$(PWD)/bin go install gotest.tools/gotestsum@latest

.PHONY: test
test: ./bin/gotestsum
	PA_POKEAPI_POKEDEX_URL="https://pokeapi.co/api/v2/pokedex" \
	./bin/gotestsum --format testname -- -tags test ./... -count=1
