package query

func Draft(v string) *SimpleQuery {
	return &SimpleQuery{value: "draft", key: "flag"}
}

func Flagged(v string) *SimpleQuery {
	return &SimpleQuery{value: "flagged", key: "flag"}
}

func New(v string) *SimpleQuery {
	return &SimpleQuery{value: "new", key: "flag"}
}

func Passed(v string) *SimpleQuery {
	return &SimpleQuery{value: "passed", key: "flag"}
}

func Replied(v string) *SimpleQuery {
	return &SimpleQuery{value: "replied", key: "flag"}
}

func Seen(v string) *SimpleQuery {
	return &SimpleQuery{value: "seen", key: "flag"}
}

func Thrashed(v string) *SimpleQuery {
	return &SimpleQuery{value: "thrashed", key: "flag"}
}

func Attach(v string) *SimpleQuery {
	return &SimpleQuery{value: "attach", key: "flag"}
}

func Signed(v string) *SimpleQuery {
	return &SimpleQuery{value: "signed", key: "flag"}
}

func Encrypted(v string) *SimpleQuery {
	return &SimpleQuery{value: "encrypted", key: "flag"}
}
