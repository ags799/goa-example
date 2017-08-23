NAME := goa-example
PACKAGE := github.com/ags799/$(NAME)
VERSION := $(shell git describe --tags --always --dirty='-dev')
GOA_GENERATED := app client swagger tool

.PHONY: all
all: test lint

.PHONY: devtools
devtools:
	go get -u github.com/FiloSottile/gvt

.PHONY: clean
clean:
	-rm -r $(NAME) $(GOA_GENERATED)

.PHONY: run
run: $(NAME)
	./$(NAME)

.PHONY: test
test: generated
	go test

.PHONY: lint
lint: generated
	golint -set_exit_status

$(NAME): generated
	go build

.PHONY: generated
generated: app client swagger tool

$(GOA_GENERATED): design/design.go
	goagen bootstrap -d $(PACKAGE)/design
