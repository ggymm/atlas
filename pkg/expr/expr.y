%{
package expr

type ExprLex struct {
    tokens []string
    pos    int
    result *CondNode
}
%}

%union {
    str  string
    node *CondNode
}

// 定义终结符
%token <str> IDENT
%token AND OR
%token LPAREN RPAREN

// 定义非终结符及其类型
%type <node> or_expr
%type <node> and_expr
%type <node> primary

// 定义优先级和结合性
%left OR
%left AND

%%

or_expr: and_expr
    | or_expr OR and_expr
    {
        $$ = &CondNode{
            Type:     NodeOr,
            Children: []*CondNode{$1, $3},
        }
        yylex.(*ExprLex).result = $$
    }
    ;

and_expr: primary
    | and_expr AND primary
    {
        $$ = &CondNode{
            Type:     NodeAnd,
            Children: []*CondNode{$1, $3},
        }
        yylex.(*ExprLex).result = $$
    }
    ;

primary: IDENT
    {
        $$ = &CondNode{
            Type: NodeCond,
            Cond: $1,
        }
        yylex.(*ExprLex).result = $$
    }
    | LPAREN or_expr RPAREN
    {
        $$ = $2
        yylex.(*ExprLex).result = $$
    }
    ;

%%