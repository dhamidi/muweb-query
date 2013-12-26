.POSIX:

test: constructors.go flags.go
	go test

flags.go: data/mu-search-flags
	PATH=$$PWD/tools:$$PATH \
	emit-flag-constructor < $< > $@

constructors.go: data/mu-search-fields
	PATH=$$PWD/tools:$$PATH \
	emit-query-constructor < $< > $@
