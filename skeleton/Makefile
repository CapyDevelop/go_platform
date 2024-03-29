include .env
export

.PHONY: run test lint lint-last install-lint

GO:=go

LINT_VERSION ?= 1.56.2

LINT_BIN ?= ./bin/golangci-lint

IS_LINT_INSTALLED := $(shell $(LINT_BIN) version 2> /dev/null | grep $(LINT_VERSION))

help: ## Display this help screen
	@echo "Makefile available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  * \033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Run the project
	go run cmd/app/main.go

test: ## Run tests
	@echo "+ $@"
	@$(GO) test -race -failfast -timeout 300s -coverprofile=.test_coverage ./... && \
	$(GO) tool cover -func=.test_coverage| tail -n1 | awk '{print "Test coverage in *_test files: " $$3}'
	@rm .test_coverage

lint: install-lint ## Lint the source files
	@echo "+ $@"
	@$(LINT_BIN) run --timeout=5m

lint-last: install-lint ## Lint files from the last commit
	@echo "+ $@"
	@$(LINT_BIN) run --config=.golangci.yml --timeout=60m --new-from-rev=origin/main --whole-files --fast -v

install-lint: ## Install linter
ifndef IS_LINT_INSTALLED
	@echo "+ $@"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v$(LINT_VERSION)
endif