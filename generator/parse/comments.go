package parse

import "go/ast"

func GetCommentGroup(decl ast.Decl) *ast.CommentGroup {
	switch decl := decl.(type) {
	case *ast.FuncDecl:
		return decl.Doc
	}
	return nil
}

func GetField(decl ast.Decl, key string) *string {
	comment := GetCommentGroup(decl)
	if comment == nil {
		return nil
	}
	func
}
