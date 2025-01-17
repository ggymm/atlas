package expr

// ParseExpr 是主入口函数
func ParseExpr(input string) *CondNode {
	l := NewLexer(input)
	yyParse(l)
	return l.result
}
