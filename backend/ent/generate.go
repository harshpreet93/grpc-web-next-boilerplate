package ent

import (
	"log"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
)

// Generate generates code for DB schemas
func Generate() {
	err := entc.Generate("./schema", &gen.Config{
		Header: "// Your Custom Header",
		IDType: &field.TypeInfo{Type: field.TypeInt},
	})
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
