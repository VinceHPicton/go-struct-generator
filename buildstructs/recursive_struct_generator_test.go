package buildstructs

import "testing"

func Test_CreateStructsForFieldSlice(t *testing.T) {

	tests := map[string]struct {
		input    []StructField
		expected string
	}{
		"struct 1": {
			input: []StructField{
				{
					Name:     "parent",
					DataType: "struct",
					Tag:      "",
					SubFields: []StructField{
						{
							Name:      "shouldbestring",
							DataType:  "string",
							Tag:       "`string`",
							SubFields: nil,
						},
						{
							Name:     "shouldbestruct",
							DataType: "struct",
							Tag:      "`astruct`",
							SubFields: []StructField{
								{
									Name:      "shouldbeint",
									DataType:  "int",
									Tag:       "`int`",
									SubFields: nil,
								},
							},
						},
					},
				},
			},
			expected: "type parent struct {\n" +
				"\tshouldbestring\tstring\t`string`\n" +
				"\tshouldbestruct\tshouldbestruct\t`astruct`\n" +
				"}\n\n" +
				"type shouldbestruct struct {\n" +
				"\tshouldbeint\tint\t`int`\n" +
				"}\n\n",
		},
		"struct 1 with  parent struct as []struct (this should not change result)": {
			input: []StructField{
				{
					Name:     "parent",
					DataType: "[]struct",
					Tag:      "",
					SubFields: []StructField{
						{
							Name:      "shouldbestring",
							DataType:  "string",
							Tag:       "`string`",
							SubFields: nil,
						},
						{
							Name:     "shouldbestruct",
							DataType: "struct",
							Tag:      "`astruct`",
							SubFields: []StructField{
								{
									Name:      "shouldbeint",
									DataType:  "int",
									Tag:       "`int`",
									SubFields: nil,
								},
							},
						},
					},
				},
			},
			expected: "type parent struct {\n" +
				"\tshouldbestring\tstring\t`string`\n" +
				"\tshouldbestruct\tshouldbestruct\t`astruct`\n" +
				"}\n\n" +
				"type shouldbestruct struct {\n" +
				"\tshouldbeint\tint\t`int`\n" +
				"}\n\n",
		},
		"struct 1 with  parent and child struct as []struct (this will change result)": {
			input: []StructField{
				{
					Name:     "parent",
					DataType: "[]struct",
					Tag:      "",
					SubFields: []StructField{
						{
							Name:      "shouldbestring",
							DataType:  "string",
							Tag:       "`string`",
							SubFields: nil,
						},
						{
							Name:     "shouldbestruct",
							DataType: "[]struct",
							Tag:      "`astruct`",
							SubFields: []StructField{
								{
									Name:      "shouldbeint",
									DataType:  "int",
									Tag:       "`int`",
									SubFields: nil,
								},
							},
						},
					},
				},
			},
			expected: "type parent struct {\n" +
				"\tshouldbestring\tstring\t`string`\n" +
				"\tshouldbestruct\t[]shouldbestruct\t`astruct`\n" +
				"}\n\n" +
				"type shouldbestruct struct {\n" +
				"\tshouldbeint\tint\t`int`\n" +
				"}\n\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {

			got := CreateStructStringsForFieldSlice(test.input)

			var compareStr string
			for _, str := range got {
				compareStr += str
			}

			if compareStr != test.expected {
				t.Errorf("got:\n%v expected:\n%v", compareStr, test.expected)
			}

		})
	}
}

func Test_createStructDefinition(t *testing.T) {

	tests := map[string]struct {
		field    StructField
		expected string
	}{
		"struct 1": {
			field: StructField{
				Name:     "parent",
				DataType: "struct",
				Tag:      "",
				SubFields: []StructField{
					{
						Name:      "shouldbeint",
						DataType:  "int",
						Tag:       "`one`",
						SubFields: nil,
					},
					{
						Name:      "shouldbeint64",
						DataType:  "int64",
						Tag:       "`two`",
						SubFields: nil,
					},
					{
						Name:      "shouldbestring",
						DataType:  "string",
						Tag:       "`3`",
						SubFields: nil,
					},
					{
						Name:     "shouldbestruct",
						DataType: "struct",
						Tag:      "`astruct`",
						SubFields: []StructField{
							{},
						},
					},
				},
			},
			expected: "type parent struct {\n" +
				"\tshouldbeint\tint\t`one`\n" +
				"\tshouldbeint64\tint64\t`two`\n" +
				"\tshouldbestring\tstring\t`3`\n" +
				"\tshouldbestruct\tshouldbestruct\t`astruct`\n" +
				"}\n\n",
		},
		"struct 1 with struct datatype as []struct (this should not change the result)": {
			field: StructField{
				Name:     "parent",
				DataType: "[]struct",
				Tag:      "",
				SubFields: []StructField{
					{
						Name:      "shouldbeint",
						DataType:  "int",
						Tag:       "`one`",
						SubFields: nil,
					},
					{
						Name:      "shouldbeint64",
						DataType:  "int64",
						Tag:       "`two`",
						SubFields: nil,
					},
					{
						Name:      "shouldbestring",
						DataType:  "string",
						Tag:       "`3`",
						SubFields: nil,
					},
					{
						Name:     "shouldbestruct",
						DataType: "struct",
						Tag:      "`astruct`",
						SubFields: []StructField{
							{},
						},
					},
				},
			},
			expected: "type parent struct {\n" +
				"\tshouldbeint\tint\t`one`\n" +
				"\tshouldbeint64\tint64\t`two`\n" +
				"\tshouldbestring\tstring\t`3`\n" +
				"\tshouldbestruct\tshouldbestruct\t`astruct`\n" +
				"}\n\n",
		},
		"struct 1 with parent and child struct datatype as []struct (this will change the result)": {
			field: StructField{
				Name:     "parent",
				DataType: "[]struct",
				Tag:      "",
				SubFields: []StructField{
					{
						Name:      "shouldbeint",
						DataType:  "int",
						Tag:       "`one`",
						SubFields: nil,
					},
					{
						Name:      "shouldbeint64",
						DataType:  "int64",
						Tag:       "`two`",
						SubFields: nil,
					},
					{
						Name:      "shouldbestring",
						DataType:  "string",
						Tag:       "`3`",
						SubFields: nil,
					},
					{
						Name:     "shouldbestruct",
						DataType: "[]struct",
						Tag:      "`astruct`",
						SubFields: []StructField{
							{},
						},
					},
				},
			},
			expected: "type parent struct {\n" +
				"\tshouldbeint\tint\t`one`\n" +
				"\tshouldbeint64\tint64\t`two`\n" +
				"\tshouldbestring\tstring\t`3`\n" +
				"\tshouldbestruct\t[]shouldbestruct\t`astruct`\n" +
				"}\n\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {

			got := createStructDefinition(test.field)

			if got != test.expected {
				t.Errorf("got:\n%v expected:\n%v", got, test.expected)
			}

		})
	}
}
