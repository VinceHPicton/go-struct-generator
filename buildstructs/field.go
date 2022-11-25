package buildstructs

import (
	"fmt"
)

// StructField is used to represent both a field of a struct, or a full struct definition (if the DataType is struct or []struct)
// In which case the SubFields attribute is populated with the StructFields of that struct
type StructField struct {
	Name      string        `json:"name"`
	DataType  string        `json:"type"`
	Tag       string        `json:"tag"`
	SubFields []StructField `json:"fields"`
}

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

// GenerateStructField returns the string for a StructField which represents that field within a Go struct definition
func (f StructField) GenerateStructField() string {

	dataTypeForFieldOfStructDef := f.DataType

	if f.IsStruct() {
		dataTypeForFieldOfStructDef = f.Name
	}
	if f.IsStructSlice() {
		dataTypeForFieldOfStructDef = fmt.Sprintf("[]%v", f.Name)
	}

	return fmt.Sprintf("\t%v\t%v\t%v\n", f.Name, dataTypeForFieldOfStructDef, f.Tag)
}
