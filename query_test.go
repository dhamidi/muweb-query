package query

import (
	"testing"
)

func TestQuery_String(t *testing.T) {
	q := Subject("foo")
	if actual, expected := q.String(), "subject:'foo'"; actual != expected {
		t.Errorf("Expected q.String() to be %#v, got %#v", expected, actual)
	}
}

func TestQuery_And(t *testing.T) {
	q := Subject("foo").And(To("bar"))
	if actual, expected := q.String(), "subject:'foo' to:'bar' "; actual != expected {
		t.Errorf("Expected q.String() to be %#v, got %#v", expected, actual)
	}
}

func TestQuery_Reverse(t *testing.T) {
	q := Subject("foo").Reverse()
	if actual, expected := q.String(), "subject:'foo' --sortfield='date' --reverse"; actual != expected {
		t.Errorf("Expected q.String to be %#v, got %#v", expected, actual)
	}

}
func TestQuery_SortBy(t *testing.T) {
	q := Subject("foo").SortBy("from")
	if actual, expected := q.String(), "subject:'foo' --sortfield='from'"; actual != expected {
		t.Errorf("Expected q.String() to be %#v, got %#v", expected, actual)
	}
}
