package api

import (
	"fmt"
	"strings"

	"atlas/pkg/expr"
)

func parseExpr(input string) *expr.CondNode {
	return expr.ParseExpr(input)
}

func buildQuery(node *expr.CondNode, args *[]any) string {
	if node == nil {
		return ""
	}

	switch node.Type {
	case expr.NodeAnd:
		// 处理 AND 运算符
		if len(node.Children) > 1 {
			conds := make([]string, 0, len(node.Children))
			for _, child := range node.Children {
				conds = append(conds, BuildQuery(child, args))
			}
			// 使用 AND 连接条件
			return fmt.Sprintf("(%s)", strings.Join(conds, " AND "))
		}
	case expr.NodeOr:
		// 处理 OR 运算符
		if len(node.Children) > 1 {
			conds := make([]string, 0, len(node.Children))
			for _, child := range node.Children {
				conds = append(conds, BuildQuery(child, args))
			}
			// 使用 OR 连接条件
			return fmt.Sprintf("(%s)", strings.Join(conds, " OR "))
		}
	case expr.NodeCond:
		// 处理 单一条件
		*args = append(*args, "%"+node.Cond+"%")
		return "title like ?"
	}
	return ""
}
