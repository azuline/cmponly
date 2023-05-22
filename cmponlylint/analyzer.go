package cmponlylint

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/azuline/cmponly/internal/slices"
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
				return false
			}

			// Get the first argument, which is an instance of the struct type
			// that we're selecting fields to compare on.
			//
			// We support two cases for the struct arg:
			// - Case 1: `StructType{}` -> CompositeLit
			// - Case 2: `variable (of type StructType)` -> Ident
			var structure *types.Struct
			if lit, ok := ce.Args[0].(*ast.CompositeLit); ok {
				if ident, ok := lit.Type.(*ast.Ident); ok {
					obj := pass.TypesInfo.ObjectOf(ident)
					if typeName, ok := obj.(*types.TypeName); ok {
						if s, ok := typeName.Type().Underlying().(*types.Struct); ok {
							structure = s
						}
					}
				}
			} else if ident, ok := ce.Args[0].(*ast.Ident); ok {
				obj := pass.TypesInfo.ObjectOf(ident)
				if varr, ok := obj.(*types.Var); ok {
					if named, ok := varr.Type().(*types.Named); ok {
						if s, ok := named.Underlying().(*types.Struct); ok {
							structure = s
						}
					}
				}
			}
			// If we can't parse out the struct, exit early.
			if structure == nil {
				return false
			}

			// And now loop over the struct and list the valid fields on the struct.
			var validFields []string
			for i := 0; i < structure.NumFields(); i++ {
				validFields = append(validFields, structure.Field(i).Name())
			}

			// Collect string representations of the fields we're selecting on
			// the structType.
			var userSpecifiedFields []string
			for _, fieldArg := range ce.Args[1:] {
				if lit, ok := fieldArg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
					// Ignore fields with errors, means the passed in fields
					// are variable. IDK why that would be the case.
					if unquotedValue, err := strconv.Unquote(lit.Value); err == nil {
						userSpecifiedFields = append(userSpecifiedFields, unquotedValue)
					}
				}
			}

			// And get the names of fields that do not exist on the struct.
			var invalidFields []string
			for _, usf := range userSpecifiedFields {
				if !slices.Contains(validFields, usf) {
					invalidFields = append(invalidFields, usf)
				}
			}

			// If any fields are invalid, report an error.
			if len(invalidFields) != 0 {
				pass.Reportf(ce.Pos(), "fields do not exist on struct: %s", strings.Join(invalidFields, ", "))
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
