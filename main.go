package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed main.go.tmpl
var tmpl string

// Given a type that implements ServeHTTP,
// this will generate a main.go file that
// reads its configuration from CONFIG_FILE
// and listens on PORT.
func main() {
	// Read input args
	typeID := os.Args[1]

	// Parse type ID
	path, name, err := parseTypeID(typeID)
	if err != nil {
		log.Fatal(err)
	}

	// Generate and print to the console
	t := template.Must(template.New("main.go").Parse(tmpl))
	err = t.Execute(os.Stdout, ServerTmpl{
		PkgPath:  path,
		PkgName:  filepath.Base(path),
		TypeName: name,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type ServerTmpl struct {
	PkgPath  string
	PkgName  string
	TypeName string
}

func parseTypeID(typeID string) (string, string, error) {
	i := strings.LastIndex(typeID, ".")
	if i == -1 {
		return "", "", fmt.Errorf("invalid type ID %s", typeID)
	}

	return typeID[:i], typeID[i+1:], nil
}
