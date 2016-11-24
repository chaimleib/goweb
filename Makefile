MY_PKGS=$(shell go list ./... | egrep -v '^(github\.com/|_)')

all: src/server/cli/server/main.go
	go install server/cli/server

clean:
	rm -f $(shell ls bin/* 2>/dev/null | egrep -v '^bin/(\.gitkeep|govendor)')
	rm -rf $(shell ls -d pkg/* | egrep -v '^pkg/\.gitkeep')

distclean: clean
	rm -rf src/github.com/
	rm -f $(shell ls bin/* 2>/dev/null | grep -v '^bin/\.gitkeep')

test:
	go test $(MY_PKGS)

qtest:
	go test --test.short $(MY_PKGS)

watch:
	bash/watch.sh

deps:
	bash/deps.sh

.PHONY: test qtest clean distclean watch deps

