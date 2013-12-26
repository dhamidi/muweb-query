package query

import (
	"regexp"
	"strings"
	"fmt"
)

var (
	SHELL_ESCAPE_RE = regexp.MustCompile("'")
)

type Query interface {
	Args() []string
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

func (self *SimpleQuery) Args() []string {
	return []string{self.String()}
}

func (self *SimpleQuery) String() string {
	return fmt.Sprintf("%s:%s",
		self.key,
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

func NewCompoundQuery() *CompoundQuery {
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

func (self *CompoundQuery) Reverse() Query {
	self.reverse = true
	if self.sortfield == "" {
		self.sortfield = "date"
	}
	return self
}

func (self *CompoundQuery) Args() []string {
	queries := []string{}
	for _, q := range self.queries {
		queries = append(queries, q.String())
	}

	queries = append(queries, self.sortArgs())
	return queries
}

func (self *CompoundQuery) String() string {

	return strings.Join(queries, " ")
}

func (self *CompoundQuery) sortArgs() []string {
	if self.sortfield == "" {
		return ""
	}

	s := []string{fmt.Sprintf("--sortfield=%s", shellescape(self.sortfield))}
	if self.reverse {
		s = append(s, "--reverse")
	}

	return s
}

func (self *CompoundQuery) sortString() string {
	return strings.Join(self.sortArgs(), " ")
}

func shellescape(s string) string {
	return fmt.Sprintf("'%s'", SHELL_ESCAPE_RE.ReplaceAllString(s, "'\\''"))
}
