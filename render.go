package main

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

type GeneratedType struct {
	Name      string
	LowerName string
}

func getRenderedPath(suffix, inputPath string) (string, error) {
	if !strings.HasSuffix(inputPath, ".go") {
		return "", fmt.Errorf("Input path %s doesn't have .go extension", inputPath)
	}
	trimmed := strings.TrimSuffix(inputPath, ".go")
	dir, file := filepath.Split(trimmed)
	return filepath.Join(dir, fmt.Sprintf("%s_%s.go", file, suffix)), nil
}

type generateTemplateData struct {
	Package string
	Types   []GeneratedType
}

func render(suffix string, w io.Writer, packageName string, types []GeneratedType) error {

	switch suffix {
	case "service":
		return serviceTemplate.Execute(w, generateTemplateData{packageName, types})
	case "instrumenting":
		return instrumentingTemplate.Execute(w, generateTemplateData{packageName, types})
	case "logging":
		return loggingTemplate.Execute(w, generateTemplateData{packageName, types})
	case "transport":
		return transportTemplate.Execute(w, generateTemplateData{packageName, types})
	case "registration":
		return registrationTemplate.Execute(w, generateTemplateData{packageName, types})
	}
	return errors.New("Unknown template")
}
