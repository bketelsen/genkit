package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func loadFile(inputPath string) (string, []GeneratedType) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Could not parse file: %s", err)
	}

	packageName := identifyPackage(f)
	if packageName == "" {
		log.Fatalf("Could not determine package name of %s", inputPath)
	}

	services := map[string]bool{}
	for _, decl := range f.Decls {
		typeName, ok := identifyServiceType(decl)
		if ok {
			services[typeName] = true
			continue
		}
	}

	types := []GeneratedType{}
	for typeName, _ := range services {
		lowerService := strings.ToLower(typeName)
		service := GeneratedType{typeName, lowerService}
		types = append(types, service)
	}

	return packageName, types
}

func identifyPackage(f *ast.File) string {
	if f.Name == nil {
		return ""
	}
	return f.Name.Name
}

func identifyServiceType(decl ast.Decl) (typeName string, match bool) {
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok {
		return
	}
	if genDecl.Doc == nil {
		return
	}

	found := false
	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "@service") {
			found = true
			break
		}
	}
	if !found {
		return
	}

	for _, spec := range genDecl.Specs {
		if typeSpec, ok := spec.(*ast.TypeSpec); ok {
			if typeSpec.Name != nil {
				typeName = typeSpec.Name.Name
				break
			}
		}
	}
	if typeName == "" {
		return
	}

	match = true
	return
}
