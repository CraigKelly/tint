BASEDIR=$(CURDIR)
TOOLDIR=$(BASEDIR)/script
HELPDIR=$(BASEDIR)/helper

BINARY=tint
SOURCES := $(shell find $(BASEDIR) -name '*.go')
TESTRESOURCES := $(shell find '$(BASEDIR)/res' -type f)
TESTED=.tested

build: $(BINARY)
$(BINARY): $(SOURCES) $(TESTED) $(VERSIONOUT)
	go build -i

install: build
	go install
	ln -f -s $(HELPDIR)/tintc $(GOPATH)/bin/tintc

clean:
	$(TOOLDIR)/clean

lint-install:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

format:
	go fmt *.go

lint: format
	go vet
	golint
	goconst .

test: $(TESTRESOURCES) $(TESTED)
$(TESTED): $(SOURCES)
	$(TOOLDIR)/test

bench:
	$(TOOLDIR)/bench

update: clean
	$(TOOLDIR)/update

.PHONY: build install clean lint-install format lint test bench update
