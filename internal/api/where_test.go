package api

import (
	"testing"
)

func Test_ParseCondition(t *testing.T) {
	text := []string{
		"A AND B",
		"A OR B",
		"A AND (B OR C)",
		"A AND (B OR C) AND D",
		"A AND (B OR C) AND (D OR E)",
		"A AND (B OR C) AND (D OR E) AND F",
		"A AND ((B AND C) OR (D AND E))",
	}
	for _, str := range text {
		cond := parseExpr(str)
		args := make([]any, 0)
		where := buildQuery(cond, &args)
		t.Logf("where: %+v, args: %+v", where, args)
	}
}
