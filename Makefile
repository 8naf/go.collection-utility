init: config-git install-deps

config-git:
	git config --local include.path ../.gitconfig

install-deps:
	npm install --no-package-lock --global @commitlint/cli @commitlint/config-conventional
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	sh -s -- -b $(shell go env GOPATH)/bin

lint-commit-message:
	commitlint --verbose --config ./.config/commitlint.config.ts --edit $(file)

lint-code:
	golangci-lint run --verbose --config ./.config/.golangci.yml --fix

test:
	go test -race -v -cover ./...

merge:
	git merge --no-ff $(branch) -m ":twisted_rightwards_arrows: merge: pull request #$(number) from $(branch)"

.PHONY: init config-git install-deps lint-commit-message lint-code test merge
