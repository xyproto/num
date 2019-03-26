.PHONY: all clean install

DESTDIR ?=
PREFIX ?= /usr

all:
	go build
	(cd cmd/frac; go build)

install:
	install -Dm755 -t "$(DESTDIR)$(PREFIX)/bin" cmd/frac/frac

clean:
	(cd cmd/frac; go clean)
	go clean
