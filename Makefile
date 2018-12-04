all: build

install:
	$(prefix)/bin/go get -u golang.org/x/lint/golint
	$(prefix)/bin/go get honnef.co/go/tools/cmd/megacheck

lint: install
	$(prefix)/bin/go fmt
	$(prefix)/bin/go vet
	$(GOPATH)/bin/golint -set_exit_status
	$(GOPATH)/bin/megacheck -unused.exit-non-zero -staticcheck.exit-non-zero

test: lint
	$(prefix)/bin/go test

build: clean test
	$(prefix)/bin/go build -o adventofcode

clean:
	@rm -f adventofcode

.PHONY: all install lint test build install
