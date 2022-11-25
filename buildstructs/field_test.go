package buildstructs

import (
	"fmt"
	"testing"
)

func TestGenerateStructField(t *testing.T) {
	tests := map[string]struct {
		structField StructField
		expected    string
	}{
		"string type structfield": {
			structField: StructField{
				Name:      "myField",
				DataType:  "string",
				Tag:       "`json:\"my_field\"`",
				SubFields: nil,
			},
			expected: "\tmyField\tstring\t`json:\"my_field\"`\n",
		},
		"int type structfield": {
			structField: StructField{
				Name:      "myIntField",
				DataType:  "int",
				Tag:       "`json:\"my_int_field\"`",
				SubFields: nil,
			},
			expected: "\tmyIntField\tint\t`json:\"my_int_field\"`\n",
		},
		"nonsense type structfield": {
			structField: StructField{
				Name:      "nonsense",
				DataType:  "kittens",
				Tag:       "",
				SubFields: nil,
			},
			expected: "\tnonsense\tkittens\t\n",
		},
		"struct type structfield": {
			structField: StructField{
				Name:     "myStructField",
				DataType: "struct",
				Tag:      "`json:\"my_struct_field\"`",
				SubFields: []StructField{
					{
						Name:     "child1",
						DataType: "int",
						Tag:      "",
					},
				},
			},
			expected: "\tmyStructField\tmyStructField\t`json:\"my_struct_field\"`\n",
		},
		"struct slice type structfield": {
			structField: StructField{
				Name:     "myStructSliceField",
				DataType: "[]struct",
				Tag:      "`json:\"my_struct_slice_field\"`",
				SubFields: []StructField{
					{
						Name:     "child1",
						DataType: "int",
						Tag:      "",
					},
				},
			},
			expected: "\tmyStructSliceField\t[]myStructSliceField\t`json:\"my_struct_slice_field\"`\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tc.structField.GenerateStructField()

			if actual != tc.expected {
				t.Errorf("expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}

func ExampleGenerateStructField_standard_field() {

	normalField := StructField{
		Name:      "myField",
		DataType:  "string",
		Tag:       "`json:\"my_field\"`",
		SubFields: nil,
	}

	fmt.Print(normalField.GenerateStructField())

	// Output:
	// myField	string	`json:"my_field"`
}

func ExampleGenerateStructField_struct_field() {

	structTypeField := StructField{
		Name:     "myStructField",
		DataType: "struct",
		Tag:      "`json:\"my_struct_field\"`",
		SubFields: []StructField{
			{
				Name:     "child1",
				DataType: "int",
				Tag:      "",
			},
		},
	}

	fmt.Print(structTypeField.GenerateStructField())

	// Output:
	// myStructField	myStructField	`json:"my_struct_field"`
}

func ExampleGenerateStructField_struct_slice_field() {

	structTypeField := StructField{
		Name:     "myStructSliceField",
		DataType: "[]struct",
		Tag:      "`json:\"my_struct_slice_field\"`",
		SubFields: []StructField{
			{
				Name:     "child1",
				DataType: "int",
				Tag:      "",
			},
		},
	}

	fmt.Print(structTypeField.GenerateStructField())

	// Output:
	// myStructSliceField	[]myStructSliceField	`json:"my_struct_slice_field"`
}
