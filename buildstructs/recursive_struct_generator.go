package buildstructs

import (
	"fmt"
	"strings"
)

func CreateStructStringsForFieldSlice(fields []StructField) []string {
	var structDefs []string

	generateStructsRecursion(fields, &structDefs)

	return structDefs
}

func generateStructsRecursion(fields []StructField, structDefinitions *[]string) {

	for _, f := range fields {
		if f.IsStruct() || f.IsStructSlice() {

			// in any order:
			// make a struct definition for the current struct (f) and add it to the list to be printed to stdout later
			newStructDefinition := createStructDefinition(f)
			*structDefinitions = append(*structDefinitions, newStructDefinition)

			// pass its subfields to this func - this func will kick off creating definitions for any further substructs
			generateStructsRecursion(f.SubFields, structDefinitions)
		}
	}

}

// createStructDefinition takes a StructField and returnes the text which defines the Go struct for it.
func createStructDefinition(f StructField) string {

	var sb strings.Builder

	sb.Write([]byte(fmt.Sprintf("type %v struct {\n", f.Name)))

	for _, fieldValueForString := range f.SubFields {
		sb.Write([]byte(fieldValueForString.GenerateStructField()))
	}

	sb.Write([]byte("}\n\n"))

	return sb.String()
}
