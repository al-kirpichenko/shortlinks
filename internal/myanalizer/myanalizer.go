package myanalizer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// OsExitCheckAnalyzer анализатор
var OsExitCheckAnalyzer = &analysis.Analyzer{
	Name: "osexitcheck",
	Doc:  "check for os.Exit in main",
	Run:  run,
}

// run выполняет проверку на прямой вызов os.Exit() в функции main пакета main
func run(pass *analysis.Pass) (interface{}, error) {

	// isMainPkg проверка пакета main (bool)
	isMainPkg := func(x *ast.File) bool {
		return x.Name.Name == "main"
	}

	// isMainPkg проверка функции main (bool)
	isMainFunc := func(x *ast.FuncDecl) bool {
		return x.Name.Name == "main"
	}

	isOsExit := func(x *ast.SelectorExpr, isMain bool) bool {
		if !isMain || x.X == nil {
			return false
		}
		call, ok := x.X.(*ast.Ident)
		if !ok {
			return false
		}
		if call.Name == "os" && x.Sel.Name == "Exit" {
			pass.Reportf(x.Pos(), "os.Exit called in main!")
			return true
		}
		return false
	}

	result := false
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.File:
				if !isMainPkg(x) { // если пакет не main - выходим
					return false
				}
			case *ast.FuncDecl:
				if !isMainFunc(x) { // если функция не main - выходим
					return false
				}
				result = true

			case *ast.SelectorExpr:
				if isOsExit(x, result) {
					return false
				}
			}
			return true
		})
	}
	return nil, nil
}
