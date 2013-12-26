package muweb

import (
	"regexp"
	"strings"
)

const (
	SHELL_ESCAPE_RE = regexp.MustCompile("'")
)

type Query interface {
	String() string
	And(Query) Query
	SortBy(string) Query
	Reverse() Query
}

type SimpleQuery struct {
	key   string
	value string
}

type CompoundQuery struct {
	sortfield string
	reverse   bool
	queries   []Query
}

func (self *SimpleQuery) String() string {
	return fmt.Sprintf("%s:%s",
		shellescape(self.key),
		shellescape(self.value),
	)
}

func (self *SimpleQuery) And(q Query) Query {
	return NewAndQuery(self, q)
}

func (self *SimpleQuery) SortBy(field string) Query {
	cq := NewCompoundQuery()
	return cq.SortBy(field).And(self)
}

func (self *SimpleQuery) Reverse() Query {
	cq := NewCompoundQuery()
	return cq.Reverse().And(self)
}

func NewAndQuery(a, b Query) Query {
	cq := NewCompoundQuery()
	return cq.And(a).And(b)
}

func NewCompoundQuery() CompoundQuery {
	return &CompoundQuery{
		sortfield: "",
		reverse:   false,
		queries:   []Query{},
	}
}

func (self *CompoundQuery) And(q Query) Query {
	self.queries = append(self.queries, q)
	return self
}

func (self *CompoundQuery) SortBy(field string) Query {
	self.sortfield = field
	return self
}

func (self *CompoundQuery) reverse() Query {
	self.reverse = true
	return self
}

func (self *CompoundQuery) String() string {
	queries = []string{}
	for q := range self.queries {
		queries = append(queries, q.String())
	}

	queries = append(queries, self.sortString())
	return strings.Join(queries, " ")
}

func (self *CompoundQuery) sortString() string {
	if self.sortfield == "" {
		return ""
	}

	s := []string{fmt.Sprintf("--sortfield=%s", shellescape(self.sortfield))}
	if self.reverse {
		s = append(s, "--reverse")
	}

	return strings.Join(s, " ")
}

func shellescape(s string) string {
	return fmt.Sprintf("'%s'", SHELL_ESCAPE_RE.ReplaceAllString(s, "'\\''"))
}