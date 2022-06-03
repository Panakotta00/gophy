package project

import (
	"context"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path"
	"strings"
)

type Controller struct {
	Route    Route
	Package  string
	FuncName string
}

func ParseControllerFile(ctx context.Context, path string) ([]*Controller, error) {
	fset := token.NewFileSet()
	absPath := AbsolutePath(ctx, path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	file, err := parser.ParseFile(fset, absPath, data, parser.Mode(0))
	if err != nil {
		return nil, err
	}

	var controller []*Controller
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			controller = append(controller, FuncToController(path, decl))
		}
	}

	return nil, nil
}

func ControllerRoute(filePath string, decl *ast.FuncDecl) *Route {
	// remove first node from `controller/my/controller/path.go` to `my/controller/path.go`
	filePath = filePath[strings.Index(filePath, "/")+1:]
	filePath = path.Dir(filePath)

	if
		// TODO: route renamer
		filePath = path.Join(filePath, funcName)
	}

	return NewRoute(filePath)
}

func FuncToController(filePath string, decl *ast.FuncDecl) *Controller {
	c := &Controller{
		Route: ControllerRoute(filePath, decl.Name),
	}
}
