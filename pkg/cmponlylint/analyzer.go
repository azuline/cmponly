package cmponlylint

import (
	"go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "cmponly",
	Doc:  "checks that all fields passed into cmponly exist",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			// Filter for function calls.
			ce, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			// Filter for calls to the cmponly.Fields function.
			if funcName(ce.Fun) != "cmponly.Fields" {
				return true
			}
			// If less than 2 args, we have nothing to check. So exit early.
			if len(ce.Args) < 2 {
				return true
			}

			// Get the first argument, which is an instance of the struct type
			// that we're selecting fields to compare on.
			{
				structTypeArg := ce.Args[0]
				lit, ok := structTypeArg.(*ast.CompositeLit)
				// If we can't parse the struct out, we can't do any comparisons,
				// so ignore this node.
				if !ok {
					return true
				}

				type_, ok := lit.Type.(*ast.Ident)
				// Same deal as before; if we can't parse out the struct type,
				// rip.
				if !ok {
					return true
				}

				pass.Reportf(ce.Pos(), "found struct type: %+v\n", type_)
				return false
			}

			// Collect string representations of the fields we're selecting on
			// the structType.
			var args []string
			for _, fieldArg := range ce.Args[1:] {
				if lit, ok := fieldArg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
					// Ignore fields with errors, means the passed in fields
					// are variable. IDK why that would be the case.
					if unquotedValue, err := strconv.Unquote(lit.Value); err != nil {
						args = append(args, unquotedValue)
					}
				}
			}

			return false
		})
	}
	return nil, nil
}

// funcName transforms a CallExpr's Fun into a `package.funcName` string for
// the function its calling.
func funcName(n ast.Expr) string {
	if sel, ok := n.(*ast.SelectorExpr); ok {
		if x, ok := sel.X.(*ast.Ident); ok {
			return x.String() + "." + sel.Sel.String()
		}
	}
	return ""
}
