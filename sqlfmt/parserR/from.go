package parserR

import (
	"github.com/kanmu/go-sqlfmt/sqlfmt/lexer"
)

type FromExpr struct {
	Values []interface{}
	Parent Expr
	SubQueryCnt int
}

func parseFrom(tokens []lexer.Token)(*FromExpr, int, error){
	var (
		expr = &FromExpr{}
		consumed = 0
		restTokens = tokens
	)

	idx := 0
	// parseのそれぞれの関数がExprとconsumeしたcntだけを返すというインターフェースはそれで良さそう
	for {
		t := restTokens[idx]

		if expr.endTType(t.Type) {
			return expr, idx, nil
		}

		switch t.Type {
		case lexer.STARTPARENTHESIS:
			parseParenthesis(restTokens)
		case lexer.FUNCTION:
			cExpr, consumed, err := parseFunction(tokens[consumed:])
			if err != nil {
				// FIXME: エラーハンドリングする
				return nil, 0, err
			}

			cExpr.Parent = expr
			expr.Values = append(expr.Values, cExpr)
			idx += consumed
		default:
			expr.Values = append(expr.Values, t)
			idx++
		}
	}
}

func (expr *FromExpr) endTType(ttype lexer.TokenType) bool{
	for _, end := range lexer.EndOfFrom{
		if ttype == end {
			return true
		}
	}
	return false
}

func (f *FromExpr) Build() string {
	return ""
}