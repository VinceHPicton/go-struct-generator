package buildstructs

import (
	"fmt"
	"strings"
)

type StructField struct {
	Name      string        `json:"name"`
	DataType  string        `json:"type"`
	Tag       string        `json:"tag"`
	SubFields []StructField `json:"fields"`
}

// if you are trying to mutate f you use a pointer
// if you are just trying to read you just use value
// if f itself is already a pointer value, then you might use a pointer receiver because...thiscould cause a problem otherwise - you would find out with unit tests
func (f StructField) IsStruct() bool {
	if f.DataType == "struct" {
		return true
	} else {
		return false
	}
}

func (f StructField) IsStructSlice() bool {
	if f.DataType == "[]struct" {
		return true
	} else {
		return false
	}
}

func (f *StructField) GenerateStructField() string {

	dataTypeForFieldOfStructDef := f.DataType

	if f.IsStruct() {
		dataTypeForFieldOfStructDef = f.Name
	}
	if f.IsStructSlice() {
		dataTypeForFieldOfStructDef = fmt.Sprintf("[]%v", f.Name)
	}

	return fmt.Sprintf("\t%v\t%v\t%v\n", f.Name, dataTypeForFieldOfStructDef, f.Tag)
}

func CreateStructsForFieldSlice(fields []StructField) []string {
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

			// pass it's subfields to this func - this func will kick off creating definitions for any further substructs
			generateStructsRecursion(f.SubFields, structDefinitions)
		}
	}

}

func createStructDefinition(f StructField) string {

	var sb strings.Builder

	sb.Write([]byte(fmt.Sprintf("type %v struct {\n", f.Name)))

	for _, fieldValueForString := range f.SubFields {
		sb.Write([]byte(fieldValueForString.GenerateStructField()))
	}

	sb.Write([]byte("}\n\n"))

	return sb.String()
}
