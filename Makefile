.POSIX:

constructors.go: data/mu-search-fields
	PATH=$$PWD/tools:$$PATH; emit-query-constructor < $< > $@
