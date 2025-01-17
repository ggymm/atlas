package expr

import (
	"fmt"
	"strings"
)

type NodeType int

// 为 NodeType 添加 String 方法
func (t NodeType) String() string {
	switch t {
	case NodeAnd:
		return "AND"
	case NodeOr:
		return "OR"
	case NodeCond:
		return "NodeCond"
	}
	return "UNKNOWN"
}

const (
	NodeAnd NodeType = iota
	NodeOr
	NodeCond
)

type CondNode struct {
	Type     NodeType    // 节点类型 (AND/OR/NOT/Cond)
	Cond     string      // 节点数据 (仅在 Type == NodeCond 时有值)
	Children []*CondNode // 子节点列表 (用于 AND/OR 嵌套)
}

func NewLexer(input string) *ExprLex {
	// 在括号前后添加空格
	input = strings.ReplaceAll(input, "(", " ( ")
	input = strings.ReplaceAll(input, ")", " ) ")

	return &ExprLex{
		tokens: strings.Fields(strings.ToUpper(input)),
		pos:    0,
	}
}

func (l *ExprLex) Lex(val *yySymType) int {
	if l.pos >= len(l.tokens) {
		return 0
	}

	token := l.tokens[l.pos]
	l.pos++

	switch token {
	case "AND":
		return AND
	case "OR":
		return OR
	case "(":
		return LPAREN
	case ")":
		return RPAREN
	default:
		val.str = token
		return IDENT
	}
}

func (l *ExprLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
