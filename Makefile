GOCMD=go
GOBUILD=$(GOCMD) build
PATH := "${CURDIR}/bin:$(PATH)"

.PHONY: gobuildcache

bin/golangci-lint:
	script/bindown install $(notdir $@)

bin/shellcheck:
	script/bindown install $(notdir $@)

bin/octo: gobuildcache
	${GOBUILD} -o $@ -ldflags "-s -w" ./

bin/goreleaser:
	script/bindown install $(notdir $@)

bin/semver-next:
	script/bindown install $(notdir $@)
